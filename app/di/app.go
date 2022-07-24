package di

import "portfolio-backend/domain"

type App struct {
	RSSFeedHandler domain.RSSFeedHandler
	BlogHandler    domain.BlogHandler
}
