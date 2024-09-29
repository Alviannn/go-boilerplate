package dependencies

import (
	"go-boilerplate/internal/domains"
	"go-boilerplate/pkg/customvalidator"
	"go-boilerplate/pkg/databases"
	"go-boilerplate/pkg/logger"

	"github.com/defval/di"
)

// New creates a new DI (dependency injection) container.
func New(extraDeps ...di.Option) (container *di.Container, err error) {
	if err = logger.SetupLogger(); err != nil {
		return
	}

	// Set logging for dependency registery and resolving.
	di.SetTracer(&logger.DITracer{})

	deps := []di.Option{
		di.Provide(customvalidator.New),
		di.Provide(databases.NewMySQLDB),
		domains.Modules,
	}
	deps = append(deps, extraDeps...)

	return di.New(deps...)
}
