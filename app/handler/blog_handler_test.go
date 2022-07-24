package handler

import (
	"encoding/json"
	mock_domain "portfolio-backend/domain/mock"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	blogs_pb "github.com/igsr5/portfolio-proto/go/lib/blogs"
)

func TestBatchGetBlogs(t *testing.T) {
	tests := []struct {
		name           string
		request        events.APIGatewayProxyRequest
		mockFn         func(mr *mock_domain.MockBlogRepository)
		wantStatusCode int
		wantBlogs      []*blogs_pb.Blog
	}{
		{
			name: "get all blogs",
			request: events.APIGatewayProxyRequest{
				HTTPMethod: "GET",
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
			mockFn: func(mr *mock_domain.MockBlogRepository) {
			},
			wantStatusCode: 200,
			wantBlogs:      []*blogs_pb.Blog{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mbr := mock_domain.NewMockBlogRepository(mockCtrl)
			tt.mockFn(mbr)

			h := NewBlogHandler(mbr)

			res, err := h.BatchGetBlogs(tt.request)
			if err != nil {
				t.Errorf("failed to BatchGetBlogs. err: %v", err)
			}

			if res.StatusCode != tt.wantStatusCode {
				t.Errorf("bad status code. got: %d, want %d", res.StatusCode, tt.wantStatusCode)
			}

			var resBody blogs_pb.BatchGetBlogsResponse
			err = json.Unmarshal([]byte(res.Body), &resBody)
			if err != nil {
				t.Fatalf("failed to unmarshal response body. %v", err)
			}

			if cmp.Equal(resBody.Blogs, tt.wantBlogs) {
				t.Fatalf("bad blogs. diff: %v", cmp.Diff(resBody.Blogs, tt.wantBlogs))
			}
		})
	}
}
