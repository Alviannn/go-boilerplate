package health

import (
	health_delivery_rest "go-boilerplate/internal/domains/health/delivery_rest"
	domains_interfaces "go-boilerplate/internal/domains/interfaces"

	"github.com/defval/di"
)

var Module = di.Options(
	di.Provide(health_delivery_rest.NewRestDelivery, di.As(new(domains_interfaces.BaseRestDelivery))),
)
