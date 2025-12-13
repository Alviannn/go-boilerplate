package repositories_mysql

import (
	"context"
	"go-boilerplate/internal/constants"

	"gorm.io/gorm"
)

type tx struct {
	RepositoryMysql
}

func NewTx(repository RepositoryMysql) Tx {
	return &tx{
		RepositoryMysql: repository,
	}
}

func (r *tx) Transaction(ctx context.Context, fc func(newCtx context.Context) error) error {
	db := r.GetDB(ctx)

	return db.Transaction(func(gormTx *gorm.DB) error {
		newCtx := context.WithValue(ctx, constants.CtxKeyGormTransaction, gormTx)
		return fc(newCtx)
	})
}
