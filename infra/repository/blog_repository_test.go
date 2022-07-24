package repository

import (
	"context"
	"database/sql"
	"portfolio-backend/domain"
	"portfolio-backend/infra/models"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestGetBlogs(t *testing.T) {
	tests := []struct {
		name        string
		existsBlogs []domain.Blog
	}{
		{
			name: "get all blogs",
			existsBlogs: []domain.Blog{
				{
					Id:           "aaa",
					Title:        "Hello World",
					PostedAt:     time.Date(2020, time.December, 10, 23, 1, 10, 0, time.Local),
					SiteUrl:      "https://example.com",
					ThumbnailUrl: "https://example.com/thumbnail.png",
					ServiceName:  "Qiita",
				},
				{
					Id:           "bbb",
					Title:        "Hello World",
					PostedAt:     time.Date(2020, time.December, 10, 23, 1, 10, 0, time.Local),
					SiteUrl:      "https://example.com",
					ThumbnailUrl: "https://example.com/thumbnail.png",
					ServiceName:  "Qiita",
				},
			},
		},
	}

	db, err := NewDB()
	if err != nil {
		t.Fatalf("failed to NewDB. err: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteAllBlogs(t, db)
			bulkInsertBlogs(t, db, tt.existsBlogs)

			r := NewBlogRepository(db)

			getBlogs, err := r.GetBlogs(context.Background())
			if err != nil {
				t.Fatalf("failed to GetBlogs. %v", err)
			}

			if !cmp.Equal(tt.existsBlogs, getBlogs) {
				t.Fatalf("bad blogs. %s", cmp.Diff(tt.existsBlogs, getBlogs))
			}
		})
	}
}

func bulkInsertBlogs(t *testing.T, db *sql.DB, blogs []domain.Blog) {
	t.Helper()
	for i, b := range blogs {
		if i%2 == 0 {
			mblog := models.BlogFromManualItem{
				ID:           b.Id,
				Title:        b.Title,
				PostedAt:     null.TimeFrom(b.PostedAt),
				SiteURL:      b.SiteUrl,
				ThumbnailURL: b.ThumbnailUrl,
				ServiceName:  b.ServiceName,
			}

			err := mblog.Insert(context.Background(), db, boil.Infer())
			if err != nil {
				t.Fatalf("failed to init db blog data. %v", err)
			}
		} else {
			rblog := models.BlogFromRSSItem{
				ID:           b.Id,
				Title:        b.Title,
				PostedAt:     null.TimeFrom(b.PostedAt),
				SiteURL:      b.SiteUrl,
				ThumbnailURL: b.ThumbnailUrl,
				ServiceName:  b.ServiceName,
			}

			err := rblog.Insert(context.Background(), db, boil.Infer())
			if err != nil {
				t.Fatalf("failed to init db blog data. %v", err)
			}
		}
	}
}

func deleteAllBlogs(t *testing.T, db *sql.DB) {
	t.Helper()

	if _, err := db.Exec("DELETE FROM blog_from_rss_items"); err != nil {
		t.Fatal(err)
	}

	if _, err := db.Exec("DELETE FROM blog_from_manual_items"); err != nil {
		t.Fatal(err)
	}
}
