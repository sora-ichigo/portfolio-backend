package main

import (
	"log"
	"os"
	"portfolio-backend/app/di"
	"portfolio-backend/lib/sentryset"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	defer sentryset.CleanUp()

	app, err := di.NewApp()
	if err != nil {
		log.Fatalf("function failed with errors: %#v", err)
		os.Exit(1)
	}

	lambda.Start(app.RSSFeedHandler.GetRSSFeeds)
}
