package accounts_repository_mysql

import (
	"context"
	"go-boilerplate/internal/constants"
	accounts_interfaces "go-boilerplate/internal/domains/accounts/interfaces"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func New(db *gorm.DB) accounts_interfaces.RepositoryMySQL {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) getDB(ctx context.Context) (gormDB *gorm.DB) {
	gormDB = r.db

	if value := ctx.Value(constants.GormTransactionKey); value != nil {
		if gormTx, ok := value.(*gorm.DB); ok {
			gormDB = gormTx
		}
	}

	return
}
