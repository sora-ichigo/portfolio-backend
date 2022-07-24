package domain

import (
	"context"
	"time"
)

type Blog struct {
	Id           string
	Title        string
	PostedAt     time.Time
	SiteUrl      string
	ThumbnailUrl string
	ServiceName  string
}

type BlogRepository interface {
	GetBlogs(context.Context) ([]Blog, error)
}

type BlogHandler interface {
}
