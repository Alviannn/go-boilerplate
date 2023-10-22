package accounts_mysql_repository

import (
	"context"
	mysql_models "go-boilerplate/internal/models/mysql"
)

func (r *repositoryImpl) IsExistByEmail(ctx context.Context, email string) bool {
	var account mysql_models.Account
	query := r.DB.WithContext(ctx).
		Select("id").
		Where("email = ?", email).
		Limit(1).
		Find(&account)

	return query.Error == nil && account.ID != 0
}
