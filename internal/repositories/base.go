package repositories

import (
	repositories_mysql "go-boilerplate/internal/repositories/mysql"

	"github.com/defval/di"
)

func Module() di.Option {
	return di.Options(
		repositories_mysql.Module(),
	)
}
