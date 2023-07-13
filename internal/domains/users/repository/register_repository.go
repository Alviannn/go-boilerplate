package users_repository

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
)

func (r *RepositoryImpl) RegisterUser(params dtos.RegisterUserReq) error {
	newUser := models.User{
		Username: params.Username,
		FullName: params.FullName,
		Email:    params.Email,
		Password: params.Password,
	}

	return r.DB.Create(&newUser).Error
}

func (r *RepositoryImpl) IsUserExistByEmail(email string) bool {
	var user models.User

	query := r.DB.Select("id").
		Where("email = ?", email).
		Limit(1).
		Find(&user)

	return query.Error == nil && user.ID != 0
}
