package accounts_service

import (
	accounts_interfaces "go-boilerplate/internal/domains/accounts/interfaces"
)

type serviceImpl struct {
	Repository accounts_interfaces.Repository
}

func NewService(repository accounts_interfaces.Repository) accounts_interfaces.Service {
	return &serviceImpl{
		Repository: repository,
	}
}
