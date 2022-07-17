//go:build tools

package tools

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/volatiletech/sqlboiler/v4"
	_ "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql"
)

//go:generate go build -v -tags 'mysql' -o=./bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate
//go:generate go build -v -o=./bin/sqlboiler github.com/volatiletech/sqlboiler/v4
//go:generate go build -v -o=./bin/sqlboiler-mysql github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql
//go:generate go build -v -o=./bin/wire github.com/google/wire/cmd/wire
//go:generate go build -v -o=./bin/mockgen github.com/golang/mock/mockgen
