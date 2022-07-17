package repository

import (
	"context"
	"database/sql"
	"portfolio-backend/domain"
	"portfolio-backend/infra/models"

	rss_feeds_pb "github.com/igsr5/portfolio-proto/go/lib/blogs/rss_feed"
	"github.com/lithammer/shortuuid/v3"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type rssFeedRepositoryImpl struct {
	db *sql.DB
}

func NewRSSFeedRepository(db *sql.DB) domain.RSSFeedRepository {
	return rssFeedRepositoryImpl{
		db: db,
	}
}

func (r rssFeedRepositoryImpl) GetRSSFeeds(ctx context.Context) ([]domain.RSSFeed, error) {
	rfs, err := models.RSSFeeds().All(ctx, r.db)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get rss_feeds")
	}

	rssFeeds := make([]domain.RSSFeed, len(rfs))

	for i, rf := range rfs {
		rssFeeds[i] = domain.RSSFeed{
			Id:  rf.ID,
			Url: rf.URL,
		}
	}

	return rssFeeds, nil
}

func (r rssFeedRepositoryImpl) GetRSSFeed(ctx context.Context, id string) (*domain.RSSFeed, error) {
	rf, err := models.RSSFeeds(models.RSSFeedWhere.ID.EQ(id)).One(ctx, r.db)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get rss_feeds")
	}

	rssFeed := &domain.RSSFeed{
		Id:  rf.ID,
		Url: rf.URL,
	}

	return rssFeed, nil
}

func (r rssFeedRepositoryImpl) CreateRSSFeed(ctx context.Context, input rss_feeds_pb.CreateRSSFeedRequest) error {
	feedUrl := input.GetUrl()
	if feedUrl == "" {
		return errors.New("url must not be blank.")
	}

	rssFeed := models.RSSFeed{
		ID:  shortuuid.New(),
		URL: feedUrl,
	}

	err := rssFeed.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return errors.Wrap(err, "failed to insert rss feed to db")
	}

	return nil
}

func (r rssFeedRepositoryImpl) DeleteRSSFeed(ctx context.Context, id string) error {
	rf, err := models.RSSFeeds(qm.Where("id = ?", id)).One(ctx, r.db)
	if err != nil {
		return errors.Wrap(err, "failed to get rss_feeds")
	}

	_, err = rf.Delete(ctx, r.db)
	if err != nil {
		return errors.Wrap(err, "failed to delete rss_feeds")
	}

	return nil
}

func (r rssFeedRepositoryImpl) IsExistsUrl(ctx context.Context, url string) (bool, error) {
	exists, err := models.RSSFeeds(models.RSSFeedWhere.URL.EQ(url)).Exists(ctx, r.db)
	if err != nil {
		return exists, errors.Wrap(err, "failed to IsExistsUrl()")
	}

	return exists, nil
}
