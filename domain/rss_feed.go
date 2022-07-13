package domain

import (
	"context"
	"portfolio-backend/infra/models"

	"github.com/aws/aws-lambda-go/events"
	rss_feeds_pb "github.com/igsr5/portfolio-proto/go/lib/blogs/rss_feed"
)

type RSSFeedRepository interface {
	GetRssFeeds(ctx context.Context) ([]*models.RSSFeed, error)
	CreateRSSFeed(ctx context.Context, input rss_feeds_pb.CreateRSSFeedRequest) error
	IsExistsUrl(ctx context.Context, url string) (bool, error)
}

type RSSFeedHandler interface {
	GetRSSFeeds(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	CreateRSSFeed(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}
