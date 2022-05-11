package di

import (
	"database/sql"
	"os"
	"portfolio-backend/infra/repository"

	"github.com/google/wire"
	"github.com/pkg/errors"
)

func ProvideDB() (*sql.DB, error) {
	dsn := os.Getenv("DSN")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sql.Open")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "db.Ping() failed")
	}

	return db, nil
}

var RepositorySet = wire.NewSet(
	ProvideDB,
	repository.NewRSSFeedRepository,
)
