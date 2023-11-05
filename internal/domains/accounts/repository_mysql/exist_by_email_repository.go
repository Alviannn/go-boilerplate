package accounts_repository_mysql

import (
	"context"
	models_mysql "go-boilerplate/internal/models/mysql"
	"go-boilerplate/pkg/customerror"
	"net/http"
)

func (r *repositoryImpl) ExistByEmail(ctx context.Context, email string) (exist bool, err error) {
	var account models_mysql.Account

	query := r.db().WithContext(ctx).
		Select("id").
		Where("email = ?", email).
		Limit(1)

	if err = query.Find(&account).Error; err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to check account existence.")
		return
	}

	exist = account.IsExist()
	return
}
