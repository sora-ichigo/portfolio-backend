package main

import (
	"os"
	"portfolio-backend/app/di"
	"portfolio-backend/lib/sentryset"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
)

func main() {
	defer sentryset.CleanUp()

	app, err := di.NewApp()
	if err != nil {
		sentry.CaptureException(errors.Wrap(err, "failed to NewApp"))
		os.Exit(1)
	}

	lambda.Start(sentryset.WithCatchErrInAPIGateway(app.RSSFeedHandler.CreateRSSFeed))
}
