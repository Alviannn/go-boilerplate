package accounts_repository_mysql

import (
	"context"
	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
)

func (r *repositoryImpl) Register(ctx context.Context, params dtos.AccountRegisterRequest) error {
	newAccount := models_mysql.Account{
		Username: params.Username,
		FullName: params.FullName,
		Email:    params.Email,
		Password: params.Password,
	}

	return r.DB.WithContext(ctx).Create(&newAccount).Error
}
