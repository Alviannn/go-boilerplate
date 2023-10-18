package accounts

import (
	accounts_delivery_rest "go-boilerplate/internal/domains/accounts/delivery_rest"
	accounts_repository "go-boilerplate/internal/domains/accounts/repository"
	accounts_service "go-boilerplate/internal/domains/accounts/service"
	domains_interfaces "go-boilerplate/internal/domains/interfaces"

	"github.com/goava/di"
)

var Module = di.Options(
	di.Provide(accounts_repository.NewRepository),
	di.Provide(accounts_service.NewService),
	di.Provide(accounts_delivery_rest.NewRestDelivery, di.As(new(domains_interfaces.BaseRestDelivery))),
)
