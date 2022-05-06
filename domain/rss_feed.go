package domain

import (
	"context"

	rss_feeds_pb "github.com/igsr5/portfolio-proto/go/lib/blogs/rss_feed"
)

type RSSFeedRepository interface {
	CreateRSSFeed(ctx context.Context, input rss_feeds_pb.CreateRSSFeedRequest) error
}
