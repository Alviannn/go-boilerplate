package dependencies

import (
	"go-boilerplate/internal/domains"
	"go-boilerplate/pkg/logger"

	"github.com/defval/di"
)

// NewForStartup creates a new DI (dependency injection) container
// during the application startup or initialization. It sets the
// DI tracer (for logging when it loads or being used).
func NewForStartup() (container *di.Container, err error) {
	if err = logger.SetupLogger(); err != nil {
		return
	}

	// Set logging for dependency registery and resolving.
	di.SetTracer(&logger.DITracer{})

	return di.New(
		appModules,
		domains.Modules,
	)
}
