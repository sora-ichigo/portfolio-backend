package handler

import (
	"context"
	"encoding/json"
	"log"
	"portfolio-backend/domain"

	"github.com/aws/aws-lambda-go/events"
	rss_feeds_pb "github.com/igsr5/portfolio-proto/go/lib/blogs/rss_feed"
)

type postRSSFeedsHandlerImpl struct {
	rssFeedRepository domain.RSSFeedRepository
}

func NewPostRSSFeedsHandler(rssFeedRepository domain.RSSFeedRepository) domain.PostRSSFeedsHandler {
	return postRSSFeedsHandlerImpl{
		rssFeedRepository: rssFeedRepository,
	}
}

func (p postRSSFeedsHandlerImpl) Invoke(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	b := struct {
		Url string `json:"url"`
	}{}

	if err := json.Unmarshal([]byte(request.Body), &b); err != nil {
		log.Printf("failed json.Unmarshal() with errors: %#v", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "failed json.Unmarshal()",
		}, nil
	}

	err := p.rssFeedRepository.CreateRSSFeed(context.Background(), rss_feeds_pb.CreateRSSFeedRequest{Url: b.Url})
	if err != nil {
		log.Printf("failed create rss feed with errors: %#v", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "failed create rss feed",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
