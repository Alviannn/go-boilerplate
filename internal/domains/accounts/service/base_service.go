package accounts_service

import (
	accounts_interfaces "go-boilerplate/internal/domains/accounts/interfaces"

	"github.com/goava/di"
)

type ServiceImpl struct {
	di.Inject

	Repository accounts_interfaces.Repository
}

func NewService(p ServiceImpl) accounts_interfaces.Service {
	return &p
}
