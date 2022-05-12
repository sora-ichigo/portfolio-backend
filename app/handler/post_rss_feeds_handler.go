package handler

import (
	"context"
	"encoding/json"
	"portfolio-backend/domain"

	"github.com/aws/aws-lambda-go/events"
	rss_feeds_pb "github.com/igsr5/portfolio-proto/go/lib/blogs/rss_feed"
	"github.com/pkg/errors"
)

type PostRSSFeedsHandler interface {
	Invoke(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

type postRSSFeedsHandlerImpl struct {
	rssFeedRepository domain.RSSFeedRepository
}

func NewPostRSSFeedsHandler(rssFeedRepository domain.RSSFeedRepository) PostRSSFeedsHandler {
	return postRSSFeedsHandlerImpl{
		rssFeedRepository: rssFeedRepository,
	}
}

func (p postRSSFeedsHandlerImpl) Invoke(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	b := struct {
		Url string `json:"url"`
	}{}

	if err := json.Unmarshal([]byte(request.Body), &b); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
		}, errors.Wrap(err, "failed json.Unmarshal()")
	}

	err := p.rssFeedRepository.CreateRSSFeed(context.Background(), rss_feeds_pb.CreateRSSFeedRequest{Url: b.Url})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, errors.Wrap(err, "failed create rss feed")
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
