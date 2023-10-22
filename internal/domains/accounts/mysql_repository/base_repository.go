package accounts_mysql_repository

import (
	accounts_interfaces "go-boilerplate/internal/domains/accounts/interfaces"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) accounts_interfaces.MySQLRepository {
	return &repositoryImpl{
		DB: db,
	}
}
