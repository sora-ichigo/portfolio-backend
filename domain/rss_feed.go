package domain

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	rss_feeds_pb "github.com/igsr5/portfolio-proto/go/lib/blogs/rss_feed"
)

type a rss_feeds_pb.RSSFeed

type RSSFeed struct {
	Id  string
	Url string
}

type RSSFeedRepository interface {
	GetRSSFeeds(ctx context.Context) ([]RSSFeed, error)
	GetRSSFeed(ctx context.Context, id string) (*RSSFeed, error)
	CreateRSSFeed(ctx context.Context, input rss_feeds_pb.CreateRSSFeedRequest) error
	DeleteRSSFeed(ctx context.Context, id string) error
	IsExistsUrl(ctx context.Context, url string) (bool, error)
}

type RSSFeedHandler interface {
	BatchGetRSSFeeds(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	GetRSSFeed(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	CreateRSSFeed(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}
