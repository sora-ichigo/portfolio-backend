package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	d := os.Getenv("DSN")
	fmt.Println("--------------")
	fmt.Println(d)
	fmt.Println("--------------")

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       d,
	}, nil
}

func main() {
	lambda.Start(handler)
}
