//go:build tools

package tools

//go:generate ./bin/sqlboiler ./bin/sqlboiler-mysql
//go:generate ./bin/wire ./app/di
