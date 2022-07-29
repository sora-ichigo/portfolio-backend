package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
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

func (b *blogHandlerImpl) GetBlog(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx := context.Background()

	blog, err := b.blogRepository.GetBlog(ctx, request.PathParameters["id"])
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return events.APIGatewayProxyResponse{
				StatusCode: 404,
				Body:       "blog not found",
			}, nil
		}

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("failed to get blog: %v", err),
		}, err
	}

	resBody := blogs_pb.GetBlogResponse{
		Blog: &blogs_pb.Blog{
			Id:           blog.Id,
			Title:        blog.Title,
			PostedAt:     timestamppb.New(blog.PostedAt),
			SiteUrl:      blog.SiteUrl,
			ThumbnailUrl: blog.ThumbnailUrl,
			ServiceName:  blog.ServiceName,
		},
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

func (b *blogHandlerImpl) CreateBlog(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx := context.Background()

	blog := blogs_pb.CreateBlogRequest{}
	err := json.Unmarshal([]byte(request.Body), &blog)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("failed to unmarshal request: %v", err),
		}, err
	}

	err = b.blogRepository.CreateBlogFromManualItem(ctx, blog)
	if err != nil {
		// TODO: handling 400 error
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("failed to create blog: %v", err),
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
