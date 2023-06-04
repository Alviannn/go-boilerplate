package registeruser

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	gorm *gorm.DB
}

func NewRepository(gorm *gorm.DB) Repository {
	return &repositoryImpl{gorm: gorm}
}

func (r *repositoryImpl) RegisterUser(params dtos.RegisterUserReq) error {
	newUser := models.User{
		Username: params.Username,
		FullName: params.FullName,
		Email:    params.Email,
		Password: params.Password,
	}

	return r.gorm.Create(&newUser).Error
}

func (r *repositoryImpl) IsUserExistByEmail(email string) bool {
	var user models.User

	query := r.gorm.Select("id").
		Where("email = ?", email).
		Limit(1).
		Find(&user)

	return query.Error == nil && user.ID != ""
}
