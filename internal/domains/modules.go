package domains

import (
	"go-boilerplate/internal/domains/users"

	"github.com/goava/di"
)

var Modules = di.Options(
	users.Module,
)
