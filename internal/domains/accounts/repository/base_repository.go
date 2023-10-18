package accounts_repository

import (
	accounts_interfaces "go-boilerplate/internal/domains/accounts/interfaces"

	"github.com/goava/di"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	di.Inject

	DB *gorm.DB
}

func NewRepository(p RepositoryImpl) accounts_interfaces.Repository {
	return &p
}
