package accounts_repository

import (
	"context"
	"go-boilerplate/internal/models"
)

func (r *RepositoryImpl) Get(ctx context.Context, accountID int64) (account models.Account, err error) {
	err = r.DB.WithContext(ctx).First(&account, "id = ?", accountID).Error
	return
}
