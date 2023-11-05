package accounts_repository_mysql

import (
	"context"
	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
	"go-boilerplate/pkg/customerror"
	"net/http"
)

func (r *repositoryImpl) GetAll(ctx context.Context, params dtos.AccountGetAllRequest) (accounts []models_mysql.Account, err error) {
	query := r.db().WithContext(ctx)

	if params.Email != "" {
		query = query.Where("email = ?", params.Email)
	}
	if params.FullName != "" {
		query = query.Where("full_name = ?", params.FullName)
	}
	if params.Username != "" {
		query = query.Where("username = ?", params.Username)
	}

	if params.Limit != 0 {
		query = query.Limit(int(params.Limit))
	}
	if params.Offset != 0 {
		query = query.Offset(params.Offset)
	}

	if err = query.Find(&accounts).Error; err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get all accounts.")
	}
	return
}
