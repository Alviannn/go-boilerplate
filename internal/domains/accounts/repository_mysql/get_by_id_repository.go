package accounts_repository_mysql

import (
	"context"
	models_mysql "go-boilerplate/internal/models/mysql"
)

func (r *repositoryImpl) GetByID(ctx context.Context, accountID int64) (account models_mysql.Account, err error) {
	err = r.DB.
		WithContext(ctx).
		Limit(1).
		Find(&account, "id = ?", accountID).
		Error
	return
}
