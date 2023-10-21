package accounts_postgres_repository

import (
	"context"
	postgres_models "go-boilerplate/internal/models/postgres"
)

func (r *postgresRepositoryImpl) IsExistByEmail(ctx context.Context, email string) bool {
	var account postgres_models.Account
	query := r.DB.WithContext(ctx).
		Select("id").
		Where("email = ?", email).
		Limit(1).
		Find(&account)

	return query.Error == nil && account.ID != 0
}
