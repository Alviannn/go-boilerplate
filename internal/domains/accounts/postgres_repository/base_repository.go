package accounts_postgres_repository

import (
	accounts_interfaces "go-boilerplate/internal/domains/accounts/interfaces"

	"gorm.io/gorm"
)

type postgresRepositoryImpl struct {
	DB *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) accounts_interfaces.PostgresRepository {
	return &postgresRepositoryImpl{
		DB: db,
	}
}
