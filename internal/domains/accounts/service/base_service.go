package accounts_service

import (
	accounts_interfaces "go-boilerplate/internal/domains/accounts/interfaces"
)

type serviceImpl struct {
	PostgresRepository accounts_interfaces.PostgresRepository
}

func NewService(postgresRepository accounts_interfaces.PostgresRepository) accounts_interfaces.Service {
	return &serviceImpl{
		PostgresRepository: postgresRepository,
	}
}
