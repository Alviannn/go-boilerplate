package accounts_service

import (
	accounts_interfaces "go-boilerplate/internal/domains/accounts/interfaces"
)

type serviceImpl struct {
	RepositoryMySQL accounts_interfaces.RepositoryMySQL
}

func New(repositoryMySQL accounts_interfaces.RepositoryMySQL) accounts_interfaces.Service {
	return &serviceImpl{
		RepositoryMySQL: repositoryMySQL,
	}
}
