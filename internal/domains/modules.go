package domains

import (
	"go-boilerplate/internal/domains/health"
	"go-boilerplate/internal/domains/users"

	"github.com/goava/di"
)

var Modules = di.Options(
	users.Module,
	health.Module,
)
