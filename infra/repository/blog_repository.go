package repository

import (
	"context"
	"database/sql"
	"portfolio-backend/domain"
	"portfolio-backend/infra/models"
	"sort"

	"github.com/pkg/errors"
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
	} else if err == sql.ErrNoRows {
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
