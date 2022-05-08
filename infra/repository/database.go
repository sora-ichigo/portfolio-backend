package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/pkg/errors"
)

func NewDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sql.Open")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "db.Ping() failed")
	}

	return db, nil
}
