package repositories_mysql

import (
	"context"
	"go-boilerplate/internal/constants"

	"gorm.io/gorm"
)

type tx struct {
	DB *gorm.DB
}

func NewTx(db *gorm.DB) Tx {
	return &tx{
		DB: db,
	}
}

func (r *tx) Transaction(ctx context.Context, fc func(newCtx context.Context) error) error {
	db := getDB(ctx, r.DB)

	return db.Transaction(func(gormTx *gorm.DB) error {
		newCtx := context.WithValue(ctx, constants.CtxKeyGormTransaction, gormTx)
		return fc(newCtx)
	})
}
