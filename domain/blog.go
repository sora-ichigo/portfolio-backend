package domain

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type Blog struct {
	Id           string
	Title        string
	PostedAt     time.Time
	SiteUrl      string
	ThumbnailUrl string
	ServiceName  string
}

type BlogRepository interface {
	GetBlogs(ctx context.Context) ([]*Blog, error)
	GetBlog(ctx context.Context, id string) (*Blog, error)
}

type BlogHandler interface {
	BatchGetBlogs(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}
