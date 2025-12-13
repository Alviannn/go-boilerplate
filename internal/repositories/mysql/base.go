package repositories_mysql

import (
	"context"
	"go-boilerplate/internal/constants"

	"github.com/defval/di"
	"gorm.io/gorm"
)

type BaseRepositoryMysql struct {
	DB *gorm.DB
}

func NewBase(db *gorm.DB) RepositoryMysql {
	return &BaseRepositoryMysql{
		DB: db,
	}
}

func (r BaseRepositoryMysql) GetDB(ctx context.Context) (mysqlDB *gorm.DB) {
	mysqlDB = r.DB
	value := ctx.Value(constants.CtxKeyGormTransaction)

	if value == nil {
		return
	}

	txMysqlDB, ok := value.(*gorm.DB)
	if !ok {
		return
	}

	mysqlDB = txMysqlDB
	return
}

func Module() di.Option {
	return di.Options(
		di.Provide(NewBase),
		di.Provide(NewTx),
		di.Provide(NewAccounts),
	)
}
