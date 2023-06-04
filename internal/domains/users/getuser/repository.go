package getuser

import (
	"go-boilerplate/internal/models"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	gorm *gorm.DB
}

func NewRepository(gorm *gorm.DB) Repository {
	return &repositoryImpl{gorm: gorm}
}

func (r *repositoryImpl) GetUser(userID string) (user models.User, err error) {
	err = r.gorm.First(&user, "id = ?", userID).Error
	return
}
