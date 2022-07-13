package handler_test

import (
	"context"
	"database/sql"
	"portfolio-backend/app/di"
	"portfolio-backend/infra/models"
	"portfolio-backend/infra/repository"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestGetRSSFeedsHeandler(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		request    events.APIGatewayProxyRequest
		statusCode int
	}{
		{
			name: "get all rss feeds",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
			statusCode: 200,
		},
	}

	app, _ := setup(t)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := app.RSSFeedHandler.GetRSSFeeds(test.request)
			if err != nil {
				t.Fatalf("failed to RSSFeedHandler.GetRSSFeeds(). %v", err)
			}

			if res.StatusCode != test.statusCode {
				t.Fatalf("bad status code by RSSFeedHandler.GetRSSFeeds(). got: %d, want: %d", res.StatusCode, test.statusCode)
			}
		})
	}
}

func TestPostRSSFeedsHeandler(t *testing.T) {
	tests := []struct {
		name            string
		url             string
		request         events.APIGatewayProxyRequest
		statusCode      int
		isCreateRssFeed bool
	}{
		{
			name: "success",
			url:  "http://example.com",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: `{ "url": "http://example.com" }`,
			},
			statusCode:      200,
			isCreateRssFeed: true,
		},
		{
			name: "bad body",
			url:  "",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: "{}",
			},

			statusCode:      400,
			isCreateRssFeed: false,
		},
	}

	app, db := setup(t)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := app.RSSFeedHandler.CreateRSSFeed(test.request)
			if err != nil {
				t.Fatalf("failed to RSSFeedHandler.CreateRSSFeed(). %v", err)
			}

			if res.StatusCode != test.statusCode {
				t.Fatalf("bad status code by RSSFeedHandler.CreateRSSFeed(). got: %d, want: %d", res.StatusCode, test.statusCode)
			}

			exists, _ := models.RSSFeeds(models.RSSFeedWhere.URL.EQ(test.url)).Exists(context.Background(), db)
			if exists != test.isCreateRssFeed {
				t.Fatalf("bad RSSFeed exists")
			}
		})
	}
}

func setup(t *testing.T) (*di.App, *sql.DB) {
	t.Helper()

	db, err := repository.NewDB()
	if err != nil {
		t.Fatalf("failed to repository.NewDB(). %v", err)
	}

	deleteAllRssFeeds(t, db)

	app, err := di.NewApp()
	if err != nil {
		t.Fatalf("failed to di.NewApp(). %v", err)
	}

	return app, db
}

func deleteAllRssFeeds(t *testing.T, db *sql.DB) {
	t.Helper()

	if _, err := db.Exec("DELETE FROM rss_feeds"); err != nil {
		t.Fatal(err)
	}
}
