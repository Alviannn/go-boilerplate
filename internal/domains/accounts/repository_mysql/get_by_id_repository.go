package accounts_repository_mysql

import (
	"context"
	"fmt"
	models_mysql "go-boilerplate/internal/models/mysql"
	"go-boilerplate/pkg/customerror"
	"net/http"
)

func (r *repositoryImpl) GetByID(ctx context.Context, accountID int64) (account models_mysql.Account, err error) {
	query := r.db().WithContext(ctx).
		Limit(1).
		Where("id = ?", accountID)

	if err = query.Find(&account).Error; err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get account.")
		return
	}
	if !account.IsExist() {
		err = customerror.New().
			WithCode(http.StatusNotFound).
			WithMessage(fmt.Sprintf("Cannot find account with ID %d", accountID))
		return
	}
	return
}
