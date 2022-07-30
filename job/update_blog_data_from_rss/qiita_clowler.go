package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/p1ass/feeder"
	"github.com/pkg/errors"
)

type qiitaResponse struct {
	CreatedAt *time.Time `json:"created_at"`
	Title     string     `json:"title"`
	URL       string     `json:"url"`
	Body      string     `json:"body"`
	ID        string     `json:"id"`
	User      *qiitaUser `json:"user"`
}

type qiitaUser struct {
	ID string `json:"id"`
}

type qiitaCrawler struct {
	URL string
}

func NewQiitaCrawler(url string) feeder.Crawler {
	return &qiitaCrawler{
		URL: url,
	}
}

func (crawler *qiitaCrawler) Crawl() ([]*feeder.Item, error) {
	apiUrl := convertRSSToApiURL(crawler.URL)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get response from qiita.")
	}

	var qiita []*qiitaResponse
	err = json.NewDecoder(resp.Body).Decode(&qiita)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode response body.")
	}

	items := []*feeder.Item{}
	for _, i := range qiita {
		items = append(items, convertQiitaToItem(i))
	}
	return items, nil
}

func convertQiitaToItem(q *qiitaResponse) *feeder.Item {

	i := &feeder.Item{
		Title:       q.Title,
		Link:        &feeder.Link{Href: q.URL},
		Created:     q.CreatedAt,
		ID:          q.ID,
		Description: q.Body,
	}

	if q.User != nil {
		i.Author = &feeder.Author{
			Name: q.User.ID,
		}
	}
	return i
}

func convertRSSToApiURL(rssUrl string) string {
	// e.g. rssUrl → https://qiita.com/igsr5/feed

	// e.g. igsr5
	userId := strings.TrimSuffix(strings.TrimPrefix(rssUrl, "https://qiita.com/"), "/feed")

	// e.g. https://qiita.com/api/v2/users/igsr5/items"
	return fmt.Sprintf("https://qiita.com/api/v2/users/%s/items", userId)
}
