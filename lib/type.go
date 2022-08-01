package lib

import "github.com/aws/aws-lambda-go/events"

type APIGatewayFunc func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
