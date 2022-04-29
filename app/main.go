package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// FIXME: エラーレスポンスがざる
// FIXME: IAM ROLE の権限げきつよ
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	dsnKey := os.Getenv("DSN")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String("ap-northeast-1")},
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := ssm.New(sess)

	// NOTE: 取得できた
	_, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(dsnKey),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("%#v", err),
			StatusCode: 500,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body: fmt.Sprintf("Hello, %v", dsnKey),
		// Body:       fmt.Sprintf("Hello, %v", *res.Parameter.Value),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
