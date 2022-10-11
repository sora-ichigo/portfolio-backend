package main

import (
	"context"
	"log"
	"os"
	"portfolio-backend/infra/models"
	"portfolio-backend/infra/repository"
	"portfolio-backend/lib/sentryset"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/otiai10/opengraph/v2"
	"github.com/p1ass/feeder"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/getsentry/sentry-go"
)

var jst *time.Location
var cld *cloudinary.Cloudinary

func main() {
	lambda.Start(handler)
}

func handler(request events.CloudWatchEvent) error {
	// 0. setup.
	defer sentryset.CleanUp()
	var err error
	cld, err = cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_SECRET"),
	)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "failed to create cloudinary"))
		return errors.Wrap(err, "failed to create cloudinary client")
	}

	ctx := context.Background()

	db, err := repository.NewDB()
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "failed to create db"))
		return errors.Wrap(err, "failed to create db")
	}
	defer db.Close()

	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "failed to load jst"))
		return errors.Wrap(err, "failed to load jst")
	}

	// 1. get rss_feed url list.
	rssFeeds, err := models.RSSFeeds().All(ctx, db)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "failed to get rss_feeds"))
		return errors.Wrap(err, "failed to get rss_feeds")
	}

	// 2. get blog_data from rss_feed.
	blogDataList := []*models.BlogFromRSSItem{}
	for _, rssFeed := range rssFeeds {
		b, err := getBlodDataFromRSSFeed(ctx, rssFeed.URL)
		if err != nil {
			sentry.CaptureException(errors.Wrap(err, "failed to get blog_data from rss_feed"))
			return errors.Wrap(err, "failed to get blog_data from rss_feed")
		}

		blogDataList = append(blogDataList, b...)
	}

	// 3. check status updated blog_data.
	currentBlogDataList, err := models.BlogFromRSSItems().All(ctx, db)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "failed to get current blog_data"))
		return errors.Wrap(err, "failed to get current blog_data")
	}

	opt := []cmp.Option{
		cmpopts.IgnoreFields(models.BlogFromRSSItem{}, "ID", "PostedAt", "ThumbnailURL"),
		cmpopts.SortSlices(func(i, j models.BlogFromRSSItem) bool {
			return i.Title < j.Title
		}),
	}
	diff := cmp.Diff(currentBlogDataList, models.BlogFromRSSItemSlice(blogDataList), opt...)
	if diff == "" {
		log.Println("blog_data is no updating")
		return nil
	} else {
		log.Printf("blog_data diff: %s", diff)
	}

	// 4. reflesh blog data.
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "failed to begin tx"))
		return errors.Wrap(err, "failed to begin tx")
	}
	defer tx.Rollback()

	_, err = models.BlogFromRSSItems().DeleteAll(ctx, tx)
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "failed to delete blog_data"))
		return errors.Wrap(err, "failed to delete blog_data")
	}

	for _, blogData := range blogDataList {
		resp, err := cld.Upload.Upload(ctx, blogData.ThumbnailURL, uploader.UploadParams{Folder: "portfolio"})
		if err != nil {
			sentry.CaptureException(errors.Wrap(err, "failed to upload image"))
			return errors.Wrap(err, "failed to upload image")
		}

		blogData.ThumbnailURL = resp.SecureURL

		err = blogData.Insert(ctx, tx, boil.Infer())
		if err != nil {
			sentry.CaptureException(errors.Wrap(err, "failed to insert blog_data"))
			return errors.Wrap(err, "failed to insert blog_data")
		}
	}

	err = tx.Commit()
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "failed to commit tx"))
		return errors.Wrap(err, "failed to commit tx")
	}

	return nil
}

func getBlodDataFromRSSFeed(ctx context.Context, url string) ([]*models.BlogFromRSSItem, error) {
	var r feeder.Crawler
	if strings.HasPrefix(url, "https://qiita.com") {
		r = NewQiitaCrawler(url)
	} else {
		r = feeder.NewRSSCrawler(url)
	}

	serviceName := GetServiceName(url)

	items, err := feeder.Crawl(r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get blog_data")
	}

	blogs := []*models.BlogFromRSSItem{}
	for _, item := range items {
		uuid, err := uuid.NewUUID()
		if err != nil {
			return nil, errors.Wrap(err, "failed to create uuid")
		}

		var thumbnailURL string
		if item.Enclosure != nil {
			thumbnailURL = item.Enclosure.URL
		} else {
			ogp, err := opengraph.Fetch(item.Link.Href)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get ogp")
			}

			thumbnailURL = ogp.Image[0].URL
		}

		blogs = append(blogs, &models.BlogFromRSSItem{
			ID:           uuid.String(),
			Title:        item.Title,
			SiteURL:      item.Link.Href,
			PostedAt:     null.TimeFrom(item.Created.In(jst)),
			ThumbnailURL: thumbnailURL,
			ServiceName:  serviceName,
		})
	}

	return blogs, nil
}

func GetServiceName(url string) string {
	var serviceName string
	switch {
	case strings.HasPrefix(url, "https://qiita.com"):
		serviceName = "Qiita"
	case strings.HasPrefix(url, "https://note.com"):
		serviceName = "note"
	case strings.HasPrefix(url, "https://zenn.dev"):
		serviceName = "Zenn"
	case strings.HasPrefix(url, "https://medium.com"):
		serviceName = "Medium"
	default:
		serviceName = "unknown"
	}

	return serviceName
}
