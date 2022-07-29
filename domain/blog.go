package domain

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/events"
	blogs_pb "github.com/igsr5/portfolio-proto/go/lib/blogs"
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
	CreateBlogFromManualItem(ctx context.Context, input blogs_pb.CreateBlogRequest) error
}

type BlogHandler interface {
	BatchGetBlogs(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
	GetBlog(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}
