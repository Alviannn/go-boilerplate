package domains

import (
	"go-boilerplate/internal/domains/accounts"
	"go-boilerplate/internal/domains/health"

	"github.com/goava/di"
)

var Modules = di.Options(
	accounts.Module,
	health.Module,
)
