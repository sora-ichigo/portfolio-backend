package di

import (
	"portfolio-backend/app/handler"

	"github.com/google/wire"
)

var HandlerSet = wire.NewSet(
	handler.NewRSSFeedHandler,
	handler.NewBlogHandler,
)
