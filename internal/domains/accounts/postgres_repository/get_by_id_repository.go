package accounts_postgres_repository

import (
	"context"
	postgres_models "go-boilerplate/internal/models/postgres"
)

func (r *postgresRepositoryImpl) GetByID(ctx context.Context, accountID int64) (account postgres_models.Account, err error) {
	err = r.DB.
		WithContext(ctx).
		Limit(1).
		Find(&account, "id = ?", accountID).
		Error
	return
}
