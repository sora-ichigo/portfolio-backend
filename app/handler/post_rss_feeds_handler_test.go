package handler_test

import (
	"context"
	"portfolio-backend/app/di"
	"portfolio-backend/infra/models"
	"portfolio-backend/infra/repository"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestPostRSSFeedsHeandler(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		request    events.APIGatewayProxyRequest
		statusCode int
		exists     bool
	}{
		{
			name: "success",
			url:  "http://example.com",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: `{"url": "http://example.com"}`,
			},
			statusCode: 200,
			exists:     true,
		},
		{
			name:       "bad body",
			url:        "",
			request:    events.APIGatewayProxyRequest{Body: ""},
			statusCode: 400,
			exists:     false,
		},
	}

	app, err := di.NewApp()
	if err != nil {
		t.Fatalf("failed to di.NewApp(). %v", err)
	}

	for _, test := range tests {
		res, err := app.PostRSSFeedsHandler.Invoke(test.request)
		if err != nil {
			t.Fatalf("failed to PostRSSFeedsHandler.Invoke(). %v", err)
		}

		if res.StatusCode != test.statusCode {
			t.Fatalf("bad status code by PostRSSFeedsHandler.Invoke(). got: %d, want: %d", res.StatusCode, test.statusCode)
		}

		db, err := repository.NewDB()
		if err != nil {
			t.Fatalf("failed to repository.NewDB(). %v", err)
		}

		exists, _ := models.RSSFeeds(models.RSSFeedWhere.URL.EQ(test.url)).Exists(context.Background(), db)
		if exists != test.exists {
			t.Fatalf("bad RSSFeed exists")
		}
	}
}
