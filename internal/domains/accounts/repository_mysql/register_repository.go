package accounts_repository_mysql

import (
	"context"
	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
	"go-boilerplate/pkg/customerror"
	"net/http"
)

func (r *repositoryImpl) Register(ctx context.Context, params dtos.AccountRegisterRequest) (err error) {
	newAccount := models_mysql.Account{
		Username: params.Username,
		FullName: params.FullName,
		Email:    params.Email,
		Password: params.Password,
	}

	query := r.getDB(ctx).
		WithContext(ctx).
		Create(&newAccount)

	if err = query.Error; err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to register new account.")
	}
	return
}
