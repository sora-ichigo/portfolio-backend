package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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

	fmt.Println("1")
	// dsn, err := config.DSN("development")
	// if err != nil {
	// 	return events.APIGatewayProxyResponse{
	// 		StatusCode: 500,
	// 	}, errors.Wrap(err, "failed get dsn")
	// }

	// fmt.Println("2")
	// _, err = repository.NewDB(dsn)
	// if err != nil {
	// 	return events.APIGatewayProxyResponse{
	// 		StatusCode: 500,
	// 	}, errors.Wrap(err, "failed connection db")
	// }

	// fmt.Println("3")
	// r := repository.NewRSSFeedRepository(db)

	// err = r.CreateRSSFeed(context.Background(), rss_feeds_pb.CreateRSSFeedRequest{Url: b.Url})
	// if err != nil {
	// 	return events.APIGatewayProxyResponse{
	// 		StatusCode: 500,
	// 	}, errors.Wrap(err, "failed create rss feed")
	// }
	// fmt.Println("4")

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
