package accounts_repository

import (
	"context"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
)

func (r *RepositoryImpl) Register(ctx context.Context, params dtos.RegisterAccountReq) error {
	newAccount := models.Account{
		Username: params.Username,
		FullName: params.FullName,
		Email:    params.Email,
		Password: params.Password,
	}

	return r.DB.WithContext(ctx).Create(&newAccount).Error
}

func (r *RepositoryImpl) IsExistByEmail(ctx context.Context, email string) bool {
	var account models.Account
	query := r.DB.WithContext(ctx).
		Select("id").
		Where("email = ?", email).
		Limit(1).
		Find(&account)

	return query.Error == nil && account.ID != 0
}
