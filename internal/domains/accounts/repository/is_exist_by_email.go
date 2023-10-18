package accounts_repository

import (
	"context"
	"go-boilerplate/internal/models"
)

func (r *RepositoryImpl) IsExistByEmail(ctx context.Context, email string) bool {
	var account models.Account
	query := r.DB.WithContext(ctx).
		Select("id").
		Where("email = ?", email).
		Limit(1).
		Find(&account)

	return query.Error == nil && account.ID != 0
}
