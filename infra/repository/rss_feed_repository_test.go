package repository

import (
	"context"
	"portfolio-server-api/config"
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

	dsn, _ := config.DSN("test")
	db, err := NewDB(dsn)
	if err != nil {
		t.Fatalf("failed to NewDB. err: %v", err)
	}

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
