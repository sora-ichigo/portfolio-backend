package handler

import (
	"context"
	"encoding/json"
	"fmt"
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
	ctx := context.Background()

	url, err := getURL(request.Body)
	if err != nil {
		log.Printf("failed json.Unmarshal() with errors: %#v", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("failed json.Unmarshal() with errors: %#v", err),
		}, nil
	} else if url == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "bad request body. json field `url` is must be specifed.",
		}, nil
	}

	exists, err := p.rssFeedRepository.IsExistsUrl(ctx, url)
	if err != nil {
		log.Printf("failed to IsExistsUrl() with errors: %#v", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("failed to check whether exists url in rss_feeds table. with errors: %#v", err),
		}, nil
	}

	if exists {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "bad request body. specifed `url` is already exists.",
		}, nil
	}

	err = p.rssFeedRepository.CreateRSSFeed(ctx, rss_feeds_pb.CreateRSSFeedRequest{Url: url})
	if err != nil {
		log.Printf("failed create rss feed with errors: %#v", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("failed create rss feed with errors: %#v", err),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func getURL(body string) (string, error) {
	b := struct {
		Url string `json:"url"`
	}{}

	err := json.Unmarshal([]byte(body), &b)

	return b.Url, err
}
