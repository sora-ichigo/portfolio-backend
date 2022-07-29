package repository

import (
	"context"
	"database/sql"
	"portfolio-backend/domain"
	"portfolio-backend/infra/models"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	blogs_pb "github.com/igsr5/portfolio-proto/go/lib/blogs"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGetBlogs(t *testing.T) {
	tests := []struct {
		name        string
		existsBlogs []*domain.Blog
	}{
		{
			name: "get all blogs",
			existsBlogs: []*domain.Blog{
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

func TestGetBlog(t *testing.T) {
	tests := []struct {
		name        string
		existsBlogs []*domain.Blog
		id          string
		want        *domain.Blog
	}{
		{
			name: "get blog from manual item",
			existsBlogs: []*domain.Blog{
				{
					Id:           "aaa",
					Title:        "Hello World",
					PostedAt:     time.Date(2020, time.December, 10, 23, 1, 10, 0, time.Local),
					SiteUrl:      "https://example.com",
					ThumbnailUrl: "https://example.com/thumbnail.png",
					ServiceName:  "Qiita",
				},
				{
					Id:       "bbb",
					Title:    "Hello World",
					PostedAt: time.Date(2020, time.December, 10, 23, 1, 10, 0, time.Local), SiteUrl: "https://example.com",
					ThumbnailUrl: "https://example.com/thumbnail.png",
					ServiceName:  "Qiita",
				},
			},
			id: "aaa",
			want: &domain.Blog{
				Id:           "aaa",
				Title:        "Hello World",
				PostedAt:     time.Date(2020, time.December, 10, 23, 1, 10, 0, time.Local),
				SiteUrl:      "https://example.com",
				ThumbnailUrl: "https://example.com/thumbnail.png",
				ServiceName:  "Qiita",
			},
		},
		{
			name: "get blog from rss item",
			existsBlogs: []*domain.Blog{
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
			id: "bbb",
			want: &domain.Blog{
				Id:           "bbb",
				Title:        "Hello World",
				PostedAt:     time.Date(2020, time.December, 10, 23, 1, 10, 0, time.Local),
				SiteUrl:      "https://example.com",
				ThumbnailUrl: "https://example.com/thumbnail.png",
				ServiceName:  "Qiita",
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

			getBlog, err := r.GetBlog(context.Background(), tt.id)
			if err != nil {
				t.Fatalf("failed to GetBlog. %v", err)
			}

			if !cmp.Equal(tt.want, getBlog) {
				t.Fatalf("bad blog. %s", cmp.Diff(tt.want, getBlog))
			}
		})
	}
}

func TestCreateBlogFromManualItem(t *testing.T) {
	tests := []struct {
		name    string
		input   blogs_pb.CreateBlogRequest
		wantErr bool
	}{
		{
			name:    "create blog from manual item",
			input:   blogs_pb.CreateBlogRequest{Title: "Hello World", PostedAt: timestamppb.New(time.Now()), SiteUrl: "https://example.com", ThumbnailUrl: "https://example.com/thumbnail.png", ServiceName: "Qiita"},
			wantErr: false,
		},
		{
			name:    "title is required",
			input:   blogs_pb.CreateBlogRequest{PostedAt: timestamppb.New(time.Now()), SiteUrl: "https://example.com", ThumbnailUrl: "https://example.com/thumbnail.png", ServiceName: "Qiita"},
			wantErr: true,
		},
		{
			name:    "posted at is required",
			input:   blogs_pb.CreateBlogRequest{Title: "Hello World", SiteUrl: "https://example.com", ThumbnailUrl: "https://example.com/thumbnail.png", ServiceName: "Qiita"},
			wantErr: true,
		},
		{
			name:    "site url is required",
			input:   blogs_pb.CreateBlogRequest{Title: "Hello World", PostedAt: timestamppb.New(time.Now()), ThumbnailUrl: "https://example.com/thumbnail.png", ServiceName: "Qiita"},
			wantErr: true,
		},
		{
			name:    "thumbnail url is required",
			input:   blogs_pb.CreateBlogRequest{Title: "Hello World", PostedAt: timestamppb.New(time.Now()), SiteUrl: "https://example.com", ServiceName: "Qiita"},
			wantErr: true,
		},
		{
			name:    "service name is required",
			input:   blogs_pb.CreateBlogRequest{Title: "Hello World", PostedAt: timestamppb.New(time.Now()), SiteUrl: "https://example.com", ThumbnailUrl: "https://example.com/thumbnail.png"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := NewDB()
			if err != nil {
				t.Fatalf("failed to NewDB. err: %v", err)
			}

			r := NewBlogRepository(db)

			err = r.CreateBlogFromManualItem(context.Background(), tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("want err, but not get err")
				}
			} else {
				if err != nil {
					t.Fatalf("failed to CreateBlogFromManualItem. %v", err)
				}
			}
		})
	}
}

func bulkInsertBlogs(t *testing.T, db *sql.DB, blogs []*domain.Blog) {
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
