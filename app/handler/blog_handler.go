package handler

import (
	"portfolio-backend/domain"

	"github.com/aws/aws-lambda-go/events"
)

type blogHandlerImpl struct {
	blogRepository domain.BlogRepository
}

func NewBlogHandler(blogRepository domain.BlogRepository) domain.BlogHandler {
	return &blogHandlerImpl{blogRepository: blogRepository}
}

func (b *blogHandlerImpl) BatchGetBlogs(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
