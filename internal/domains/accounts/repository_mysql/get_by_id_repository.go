package accounts_repository_mysql

import (
	"context"
	mysql_models "go-boilerplate/internal/models/mysql"
)

func (r *repositoryImpl) GetByID(ctx context.Context, accountID int64) (account mysql_models.Account, err error) {
	err = r.DB.
		WithContext(ctx).
		Limit(1).
		Find(&account, "id = ?", accountID).
		Error
	return
}
