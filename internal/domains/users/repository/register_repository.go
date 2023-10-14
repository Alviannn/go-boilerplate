package users_repository

import (
	"context"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
)

func (r *RepositoryImpl) RegisterUser(ctx context.Context, params dtos.RegisterUserReq) error {
	newUser := models.User{
		Username: params.Username,
		FullName: params.FullName,
		Email:    params.Email,
		Password: params.Password,
	}

	return r.DB.WithContext(ctx).Create(&newUser).Error
}

func (r *RepositoryImpl) IsUserExistByEmail(ctx context.Context, email string) bool {
	var user models.User

	query := r.DB.WithContext(ctx).
		Select("id").
		Where("email = ?", email).
		Limit(1).
		Find(&user)

	return query.Error == nil && user.ID != 0
}
