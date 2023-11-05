package accounts_repository_mysql

import (
	accounts_interfaces "go-boilerplate/internal/domains/accounts/interfaces"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	DB *gorm.DB
	Tx *gorm.DB
}

func New(db *gorm.DB) accounts_interfaces.RepositoryMySQL {
	return &repositoryImpl{
		DB: db,
	}
}

func (r *repositoryImpl) db() *gorm.DB {
	if r.Tx != nil {
		return r.Tx
	}
	return r.db()
}
