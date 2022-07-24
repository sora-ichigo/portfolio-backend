package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"portfolio-backend/domain"

	"github.com/aws/aws-lambda-go/events"
	blogs_pb "github.com/igsr5/portfolio-proto/go/lib/blogs"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type blogHandlerImpl struct {
	blogRepository domain.BlogRepository
}

func NewBlogHandler(blogRepository domain.BlogRepository) domain.BlogHandler {
	return &blogHandlerImpl{blogRepository: blogRepository}
}

func (b *blogHandlerImpl) BatchGetBlogs(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx := context.Background()

	blogs, err := b.blogRepository.GetBlogs(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("failed to get blogs: %v", err),
		}, err
	}

	resBody := blogs_pb.BatchGetBlogsResponse{
		Blogs: make([]*blogs_pb.Blog, 0, len(blogs)),
	}

	for _, blog := range blogs {
		resBody.Blogs = append(resBody.Blogs, &blogs_pb.Blog{
			Id:           blog.Id,
			Title:        blog.Title,
			PostedAt:     timestamppb.New(blog.PostedAt),
			SiteUrl:      blog.SiteUrl,
			ThumbnailUrl: blog.ThumbnailUrl,
			ServiceName:  blog.ServiceName,
		})
	}

	resBodyStr, err := json.Marshal(resBody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("failed to marshal response: %v", err),
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(resBodyStr),
	}, nil
}
