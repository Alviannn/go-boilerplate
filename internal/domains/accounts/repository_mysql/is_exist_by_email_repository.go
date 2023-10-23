package accounts_repository_mysql

import (
	"context"
	models_mysql "go-boilerplate/internal/models/mysql"
)

func (r *repositoryImpl) IsExistByEmail(ctx context.Context, email string) (exist bool, err error) {
	var account models_mysql.Account
	err = r.DB.WithContext(ctx).
		Select("id").
		Where("email = ?", email).
		Limit(1).
		Find(&account).
		Error

	exist = account.IsExist()
	return
}
