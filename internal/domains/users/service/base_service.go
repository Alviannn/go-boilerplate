package users_service

import (
	"github.com/goava/di"
	users_interfaces "go-boilerplate/internal/domains/users/interfaces"
)

type ServiceImpl struct {
	di.Inject

	Repository users_interfaces.Repository
}

func NewService(p ServiceImpl) users_interfaces.Service {
	return &p
}
