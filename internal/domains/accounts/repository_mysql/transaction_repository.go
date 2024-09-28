package accounts_repository_mysql

import (
	"context"
	"go-boilerplate/internal/constants"

	"gorm.io/gorm"
)

func (r *repositoryImpl) Transaction(ctx context.Context, fc func(newCtx context.Context) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		newCtx := context.WithValue(ctx, constants.GormTransactionKey, tx)
		return fc(newCtx)
	})
}
