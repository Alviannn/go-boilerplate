package users_repository

import (
	"github.com/goava/di"
	"gorm.io/gorm"
	users_interfaces "go-boilerplate/internal/domains/users/interfaces"
)

type RepositoryImpl struct {
	di.Inject

	DB *gorm.DB
}

func NewRepository(p RepositoryImpl) users_interfaces.Repository {
	return &p
}
