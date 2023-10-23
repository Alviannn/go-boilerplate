package accounts_repository_mysql

import (
	"context"
	"go-boilerplate/internal/dtos"
	mysql_models "go-boilerplate/internal/models/mysql"
)

func (r *repositoryImpl) Register(ctx context.Context, params dtos.RegisterAccountReq) error {
	newAccount := mysql_models.Account{
		Username: params.Username,
		FullName: params.FullName,
		Email:    params.Email,
		Password: params.Password,
	}

	return r.DB.WithContext(ctx).Create(&newAccount).Error
}
