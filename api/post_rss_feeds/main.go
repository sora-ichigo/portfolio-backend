package main

import (
	"log"
	"os"
	"portfolio-backend/app/di"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	app, err := di.NewApp()
	if err != nil {
		log.Fatalf("function failed with errors: %v", err)
		os.Exit(1)
	}

	lambda.Start(app.PostRSSFeedsHandler)
}
