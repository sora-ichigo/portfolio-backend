package repository

import (
	"context"
	"database/sql"
	"testing"

	rss_feeds_pb "github.com/igsr5/portfolio-proto/go/lib/blogs/rss_feed"
)

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

func deleteAllRssFeeds(t *testing.T, db *sql.DB) {
	t.Helper()

	if _, err := db.Exec("DELETE FROM rss_feeds"); err != nil {
		t.Fatal(err)
	}
}
