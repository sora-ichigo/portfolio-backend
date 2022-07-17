//go:build tools

package tools

//go:generate ./bin/sqlboiler ./bin/sqlboiler-mysql
//go:generate ./bin/wire ./app/di
//go:generate ./bin/mockgen -source=./domain/rss_feed.go -destination=./domain/mock/rss_feed_mock.go
