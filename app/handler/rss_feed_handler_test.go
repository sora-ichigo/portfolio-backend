package handler

import (
	"encoding/json"
	"portfolio-backend/domain"
	mock_domain "portfolio-backend/domain/mock"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/mock/gomock"
	rss_feeds_pb "github.com/igsr5/portfolio-proto/go/lib/blogs/rss_feed"
)

func TestGetRSSFeedsHeandler(t *testing.T) {
	tests := []struct {
		name         string
		request      events.APIGatewayProxyRequest
		mockFn       func(mr *mock_domain.MockRSSFeedRepository)
		statusCode   int
		wantRssFeeds []*rss_feeds_pb.RSSFeed
	}{
		{
			name: "get all rss feeds",
			request: events.APIGatewayProxyRequest{
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
			mockFn: func(mr *mock_domain.MockRSSFeedRepository) {
				mr.EXPECT().GetRSSFeeds(gomock.Any()).Return([]domain.RSSFeed{
					{
						Id:  "aaa",
						Url: "http://example.com/1",
					},
					{
						Id:  "bbb",
						Url: "http://example.com/2",
					},
				}, nil)
			},
			statusCode: 200,
			wantRssFeeds: []*rss_feeds_pb.RSSFeed{
				{
					Id:  "aaa",
					Url: "http://example.com/1",
				},
				{
					Id:  "bbb",
					Url: "http://example.com/2",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			h, mr := setup(t, mockCtrl)
			test.mockFn(mr)

			res, err := h.GetRSSFeeds(test.request)
			if err != nil {
				t.Fatalf("failed to RSSFeedHandler.GetRSSFeeds(). %v", err)
			}

			if res.StatusCode != test.statusCode {
				t.Fatalf("bad status code by RSSFeedHandler.GetRSSFeeds(). got: %d, want: %d", res.StatusCode, test.statusCode)
			}

			var resBody rss_feeds_pb.BatchGetRSSFeedsResponse
			err = json.Unmarshal([]byte(res.Body), &resBody)
			if err != nil {
				t.Fatalf("failed to unmarshal response body. %v", err)
			}

			if !reflect.DeepEqual(resBody.RssFeeds, test.wantRssFeeds) {
				t.Fatalf("bad response body. got: %v, want: %v", resBody.RssFeeds, test.wantRssFeeds)
			}
		})
	}
}

func TestPostRSSFeedsHeandler(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		request    events.APIGatewayProxyRequest
		mockFn     func(mr *mock_domain.MockRSSFeedRepository)
		statusCode int
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
			mockFn: func(mr *mock_domain.MockRSSFeedRepository) {
				mr.EXPECT().CreateRSSFeed(gomock.Any(), gomock.Any()).Return(nil)
				mr.EXPECT().IsExistsUrl(gomock.Any(), "http://example.com").Return(false, nil)
			},
			statusCode: 200,
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
			mockFn:     func(mr *mock_domain.MockRSSFeedRepository) {},
			statusCode: 400,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			h, mr := setup(t, mockCtrl)
			test.mockFn(mr)

			res, err := h.CreateRSSFeed(test.request)
			if err != nil {
				t.Fatalf("failed to RSSFeedHandler.CreateRSSFeed(). %v", err)
			}

			if res.StatusCode != test.statusCode {
				t.Fatalf("bad status code by RSSFeedHandler.CreateRSSFeed(). got: %d, want: %d", res.StatusCode, test.statusCode)
			}
		})
	}
}

func setup(t *testing.T, mockCtrl *gomock.Controller) (domain.RSSFeedHandler, *mock_domain.MockRSSFeedRepository) {
	t.Helper()
	mockRSSFeedRepository := mock_domain.NewMockRSSFeedRepository(mockCtrl)
	return NewRSSFeedHandler(mockRSSFeedRepository), mockRSSFeedRepository
}
