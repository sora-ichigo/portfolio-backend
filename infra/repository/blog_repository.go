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

func (b *blogRepositoryImpl) GetBlogs(ctx context.Context) ([]domain.Blog, error) {
	mblog, err := models.BlogFromManualItems().All(ctx, b.db)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get BlogFromManualItems")
	}

	rblog, err := models.BlogFromRSSItems().All(ctx, b.db)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get BlogFromRSSItems")
	}

	blogs := make([]domain.Blog, 0, len(mblog)+len(rblog))
	for _, m := range mblog {
		blogs = append(blogs, domain.Blog{
			Id:           m.ID,
			Title:        m.Title,
			PostedAt:     m.PostedAt.Time,
			SiteUrl:      m.SiteURL,
			ThumbnailUrl: m.ThumbnailURL,
			ServiceName:  m.ServiceName,
		})
	}
	for _, r := range rblog {
		blogs = append(blogs, domain.Blog{
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
