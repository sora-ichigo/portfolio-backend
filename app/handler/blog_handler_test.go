package handler

import (
	mock_domain "portfolio-backend/domain/mock"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/mock/gomock"
)

func TestBatchGetBlogs(t *testing.T) {
	tests := []struct {
		name           string
		request        events.APIGatewayProxyRequest
		wantStatusCode int
	}{
		{
			name: "get all blogs",
			request: events.APIGatewayProxyRequest{
				HTTPMethod: "GET",
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
			wantStatusCode: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mbr := mock_domain.NewMockBlogRepository(mockCtrl)
			h := NewBlogHandler(mbr)

			res, err := h.BatchGetBlogs(tt.request)
			if err != nil {
				t.Errorf("failed to BatchGetBlogs. err: %v", err)
			}

			if res.StatusCode != tt.wantStatusCode {
				t.Errorf("bad status code. got: %d, want %d", res.StatusCode, tt.wantStatusCode)
			}
		})
	}
}
