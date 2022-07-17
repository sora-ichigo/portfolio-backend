package repository

import (
	"context"
	"database/sql"
	"portfolio-backend/domain"
	"portfolio-backend/infra/models"
	"reflect"
	"testing"

	rss_feeds_pb "github.com/igsr5/portfolio-proto/go/lib/blogs/rss_feed"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestGetRSSFeeds(t *testing.T) {
	tests := []struct {
		name           string
		existsRSSFeeds []domain.RSSFeed
	}{
		{
			name: "get all rss feeds",
			existsRSSFeeds: []domain.RSSFeed{
				{
					Id:  "aaa",
					Url: "http://example.com/img/1",
				},
				{
					Id:  "bbb",
					Url: "http://example.com/img/2",
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
			deleteAllRssFeeds(t, db)
			bulkInsertRssFeeds(t, db, tt.existsRSSFeeds)

			r := NewRSSFeedRepository(db)

			getRSSFeeds, err := r.GetRSSFeeds(context.Background())
			if err != nil {
				t.Fatalf("failed to GetRssFeeds. %v", err)
			}

			if !reflect.DeepEqual(getRSSFeeds, tt.existsRSSFeeds) {
				t.Fatalf("bad value returns GetRssFeeds(). want: %v, get: %v", tt.existsRSSFeeds, getRSSFeeds)
			}
		})
	}
}

func TestGetRSSFeed(t *testing.T) {
	tests := []struct {
		name           string
		targetId       string
		existsRSSFeeds []domain.RSSFeed
		wantRSSFeed    *domain.RSSFeed
	}{
		{
			name:     "get all rss feeds",
			targetId: "aaa",
			existsRSSFeeds: []domain.RSSFeed{
				{
					Id:  "aaa",
					Url: "http://example.com/img/1",
				},
				{
					Id:  "bbb",
					Url: "http://example.com/img/2",
				},
			},
			wantRSSFeed: &domain.RSSFeed{
				Id:  "aaa",
				Url: "http://example.com/img/1",
			},
		},
	}

	db, err := NewDB()
	if err != nil {
		t.Fatalf("failed to NewDB. err: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteAllRssFeeds(t, db)
			bulkInsertRssFeeds(t, db, tt.existsRSSFeeds)

			r := NewRSSFeedRepository(db)

			getRSSFeed, err := r.GetRSSFeed(context.Background(), tt.targetId)
			if err != nil {
				t.Fatalf("failed to GetRssFeed. %v", err)
			}

			if !reflect.DeepEqual(getRSSFeed, tt.wantRSSFeed) {
				t.Fatalf("bad value returns GetRssFeed(). want: %v, get: %v", tt.wantRSSFeed, getRSSFeed)
			}
		})
	}
}

func TestCreateRSSFeed(t *testing.T) {
	tests := []struct {
		name    string
		input   rss_feeds_pb.CreateRSSFeedRequest
		wantErr bool
	}{
		{
			name:    "currect pattern. can create rss_feed",
			input:   rss_feeds_pb.CreateRSSFeedRequest{Url: "https://example.com/feeds"},
			wantErr: false,
		},
		{
			name:    "bad pattern. cannot create rss_feed when no url input",
			input:   rss_feeds_pb.CreateRSSFeedRequest{},
			wantErr: true,
		},
	}

	db, err := NewDB()
	if err != nil {
		t.Fatalf("failed to NewDB. err: %v", err)
	}

	deleteAllRssFeeds(t, db)

	r := NewRSSFeedRepository(db)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.CreateRSSFeed(context.Background(), tt.input)

			if tt.wantErr && err == nil {
				t.Fatalf("bad create rss_feed. must happen error.")
			} else if !tt.wantErr && err != nil {
				t.Fatalf("failed to create rss_feed. err: %v", err)
			}
		})
	}
}

func TestIsExistsUrl(t *testing.T) {
	tests := []struct {
		name  string
		input rss_feeds_pb.CreateRSSFeedRequest
		url   string
		want  bool
	}{
		{
			name:  "exists url",
			input: rss_feeds_pb.CreateRSSFeedRequest{Url: "https://example.com/feeds/1"},
			url:   "https://example.com/feeds/1",
			want:  true,
		},
		{
			name:  "not exists url",
			input: rss_feeds_pb.CreateRSSFeedRequest{Url: "https://example.com/feeds/2"},
			url:   "https://example.com/feeds/3",
			want:  false,
		},
	}

	db, err := NewDB()
	if err != nil {
		t.Fatalf("failed to NewDB. err: %v", err)
	}

	r := NewRSSFeedRepository(db)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := r.CreateRSSFeed(context.Background(), tt.input)

			if err != nil {
				t.Fatalf("failed to create rss_feed. err: %v", err)
			}

			exists, err := r.IsExistsUrl(context.Background(), tt.url)
			if err != nil {
				t.Fatalf("failed to IsExistsUrl(). err: %v", err)
			}

			if exists != tt.want {
				t.Fatalf("bad exists. got: %v, want: %v", exists, tt.want)
			}
		})
	}

}
func TestDeleteRSSFeed(t *testing.T) {
	tests := []struct {
		name           string
		targetId       string
		existsRSSFeeds []domain.RSSFeed
	}{
		{
			name:     "get all rss feeds",
			targetId: "aaa",
			existsRSSFeeds: []domain.RSSFeed{
				{
					Id:  "aaa",
					Url: "http://example.com/img/1",
				},
				{
					Id:  "bbb",
					Url: "http://example.com/img/2",
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
			deleteAllRssFeeds(t, db)
			bulkInsertRssFeeds(t, db, tt.existsRSSFeeds)

			r := NewRSSFeedRepository(db)

			err := r.DeleteRSSFeed(context.Background(), tt.targetId)
			if err != nil {
				t.Fatalf("failed to GetRssFeed. %v", err)
			}

			exists, err := models.RSSFeeds(models.RSSFeedWhere.URL.EQ(tt.targetId)).Exists(context.Background(), db)
			if err != nil {
				t.Fatalf("failed to Exists(). err: %v", err)
			}

			if exists {
				t.Fatalf("failed to delete rss_feed. rss_feed is not deleted.")
			}
		})
	}
}

func bulkInsertRssFeeds(t *testing.T, db *sql.DB, rssFeeds []domain.RSSFeed) {
	t.Helper()
	for _, rf := range rssFeeds {
		rssFeed := models.RSSFeed{
			ID:  rf.Id,
			URL: rf.Url,
		}

		err := rssFeed.Insert(context.Background(), db, boil.Infer())
		if err != nil {
			t.Fatalf("failed to init db rss_feeds data. %v", err)
		}
	}

}

func deleteAllRssFeeds(t *testing.T, db *sql.DB) {
	t.Helper()

	if _, err := db.Exec("DELETE FROM rss_feeds"); err != nil {
		t.Fatal(err)
	}
}
