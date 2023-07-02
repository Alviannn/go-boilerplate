package users

import (
	domains_interfaces "go-boilerplate/internal/domains/interfaces"
	users_delivery_rest "go-boilerplate/internal/domains/users/delivery_rest"
	users_repository "go-boilerplate/internal/domains/users/repository"
	users_service "go-boilerplate/internal/domains/users/service"

	"github.com/goava/di"
)

var Module = di.Options(
	di.Provide(users_repository.NewRepository),
	di.Provide(users_service.NewService),
	di.Provide(users_delivery_rest.NewRestDelivery, di.As(new(domains_interfaces.BaseRestDelivery))),
)
