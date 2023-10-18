package accounts_postgres_repository

import (
	"context"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
)

func (r *postgresRepositoryImpl) Register(ctx context.Context, params dtos.RegisterAccountReq) error {
	newAccount := models.Account{
		Username: params.Username,
		FullName: params.FullName,
		Email:    params.Email,
		Password: params.Password,
	}

	return r.DB.WithContext(ctx).Create(&newAccount).Error
}
