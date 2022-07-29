package repository

import (
	"context"
	"database/sql"
	"portfolio-backend/domain"
	"portfolio-backend/infra/models"
	"sort"

	"github.com/gofrs/uuid"
	blogs_pb "github.com/igsr5/portfolio-proto/go/lib/blogs"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type blogRepositoryImpl struct {
	db *sql.DB
}

func NewBlogRepository(db *sql.DB) domain.BlogRepository {
	return &blogRepositoryImpl{db: db}
}

func (b *blogRepositoryImpl) GetBlogs(ctx context.Context) ([]*domain.Blog, error) {
	mblog, err := models.BlogFromManualItems().All(ctx, b.db)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get BlogFromManualItems")
	}

	rblog, err := models.BlogFromRSSItems().All(ctx, b.db)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get BlogFromRSSItems")
	}

	blogs := make([]*domain.Blog, 0, len(mblog)+len(rblog))
	for _, m := range mblog {
		blogs = append(blogs, &domain.Blog{
			Id:           m.ID,
			Title:        m.Title,
			PostedAt:     m.PostedAt.Time,
			SiteUrl:      m.SiteURL,
			ThumbnailUrl: m.ThumbnailURL,
			ServiceName:  m.ServiceName,
		})
	}
	for _, r := range rblog {
		blogs = append(blogs, &domain.Blog{
			Id:           r.ID,
			Title:        r.Title,
			PostedAt:     r.PostedAt.Time,
			SiteUrl:      r.SiteURL,
			ThumbnailUrl: r.ThumbnailURL,
			ServiceName:  r.ServiceName,
		})
	}

	sort.Slice(blogs, func(i, j int) bool {
		return blogs[i].PostedAt.After(blogs[j].PostedAt)
	})

	return blogs, nil
}

func (b *blogRepositoryImpl) GetBlog(ctx context.Context, id string) (*domain.Blog, error) {
	mblog, err := models.BlogFromManualItems(models.BlogFromManualItemWhere.ID.EQ(id)).One(ctx, b.db)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "failed to get BlogFromManualItems")
	} else if errors.Is(err, sql.ErrNoRows) {
		// noop
	} else {
		return &domain.Blog{
			Id:           mblog.ID,
			Title:        mblog.Title,
			PostedAt:     mblog.PostedAt.Time,
			SiteUrl:      mblog.SiteURL,
			ThumbnailUrl: mblog.ThumbnailURL,
			ServiceName:  mblog.ServiceName,
		}, nil
	}

	rblog, err := models.BlogFromRSSItems(models.BlogFromRSSItemWhere.ID.EQ(id)).One(ctx, b.db)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get BlogFromManualItems")
	}

	return &domain.Blog{
		Id:           rblog.ID,
		Title:        rblog.Title,
		PostedAt:     rblog.PostedAt.Time,
		SiteUrl:      rblog.SiteURL,
		ThumbnailUrl: rblog.ThumbnailURL,
		ServiceName:  rblog.ServiceName,
	}, nil
}

func (b *blogRepositoryImpl) CreateBlogFromManualItem(ctx context.Context, input blogs_pb.CreateBlogRequest) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return errors.Wrap(err, "failed to create uuid")
	}

	if input.Title == "" {
		return errors.New("title is required")
	}

	if input.PostedAt == nil {
		return errors.New("posted at is required")
	}

	if input.SiteUrl == "" {
		return errors.New("site url is required")
	}

	if input.ThumbnailUrl == "" {
		return errors.New("thumbnail url is required")
	}

	if input.ServiceName == "" {
		return errors.New("service name is required")
	}

	mblog := models.BlogFromManualItem{
		ID:           uuid.String(),
		Title:        input.GetTitle(),
		PostedAt:     null.TimeFrom(input.GetPostedAt().AsTime()),
		SiteURL:      input.GetSiteUrl(),
		ThumbnailURL: input.GetThumbnailUrl(),
		ServiceName:  input.GetServiceName(),
	}

	err = mblog.Insert(ctx, b.db, boil.Infer())
	if err != nil {
		return errors.Wrap(err, "failed to insert BlogFromManualItem")
	}

	return nil
}
