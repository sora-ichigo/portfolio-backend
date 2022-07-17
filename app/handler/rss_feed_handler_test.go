package handler

import (
	"database/sql"
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

			res, err := h.BatchGetRSSFeeds(test.request)
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

func TestGetRSSFeedHeandler(t *testing.T) {
	tests := []struct {
		name        string
		request     events.APIGatewayProxyRequest
		mockFn      func(mr *mock_domain.MockRSSFeedRepository)
		statusCode  int
		wantRssFeed *rss_feeds_pb.RSSFeed
	}{
		{
			name: "get specified rss_feed",
			request: events.APIGatewayProxyRequest{
				Path: "/rss_feeds/aaa",
				PathParameters: map[string]string{
					"id": "aaa",
				},
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
			mockFn: func(mr *mock_domain.MockRSSFeedRepository) {
				mr.EXPECT().GetRSSFeed(gomock.Any(), "aaa").Return(&domain.RSSFeed{
					Id:  "aaa",
					Url: "http://example.com/1",
				}, nil)
			},
			statusCode: 200,
			wantRssFeed: &rss_feeds_pb.RSSFeed{
				Id:  "aaa",
				Url: "http://example.com/1",
			},
		},
		{
			name: "not found rss_feed",
			request: events.APIGatewayProxyRequest{
				Path: "/rss_feeds/ccc",
				PathParameters: map[string]string{
					"id": "ccc",
				},
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
			mockFn: func(mr *mock_domain.MockRSSFeedRepository) {
				mr.EXPECT().GetRSSFeed(gomock.Any(), "ccc").Return(nil, sql.ErrNoRows)
			},
			statusCode:  404,
			wantRssFeed: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			h, mr := setup(t, mockCtrl)
			test.mockFn(mr)

			res, err := h.GetRSSFeed(test.request)
			if err != nil {
				t.Fatalf("failed to RSSFeedHandler.GetRSSFeed(). %v", err)
			}

			if res.StatusCode != test.statusCode {
				t.Fatalf("bad status code by RSSFeedHandler.GetRSSFeed(). got: %d, want: %d", res.StatusCode, test.statusCode)
			}

			if res.StatusCode == 200 {
				var resBody rss_feeds_pb.GetRSSFeedResponse
				err = json.Unmarshal([]byte(res.Body), &resBody)
				if err != nil {
					t.Fatalf("failed to unmarshal response body. %v", err)
				}

				if !reflect.DeepEqual(resBody.RssFeed, test.wantRssFeed) {
					t.Fatalf("bad response body. got: %v, want: %v", resBody.RssFeed, test.wantRssFeed)
				}
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
