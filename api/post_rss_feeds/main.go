package main

import (
	"context"
	"encoding/json"
	"portfolio-server-api/config"
	"portfolio-server-api/infra/repository"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	rss_feeds_pb "github.com/igsr5/portfolio-proto/go/lib/blogs/rss_feed"
	"github.com/pkg/errors"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	b := struct {
		Url string `json:"url"`
	}{}

	if err := json.Unmarshal([]byte(request.Body), &b); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, errors.Wrap(err, "failed json.Unmarshal()")
	}

	// NOTE: 今はローカルのみで動く
	// TODO: wire でいい感じにする
	dsn, err := config.DSN("development")
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, errors.Wrap(err, "failed get dsn")
	}

	db, err := repository.NewDB(dsn)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, errors.Wrap(err, "failed connection db")
	}

	r := repository.NewRSSFeedRepository(db)

	err = r.CreateRSSFeed(context.Background(), rss_feeds_pb.CreateRSSFeedRequest{Url: b.Url})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, errors.Wrap(err, "failed create rss feed")
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
