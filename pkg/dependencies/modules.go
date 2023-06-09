package dependencies

import (
	"go-boilerplate/pkg/customvalidator"
	"go-boilerplate/pkg/databases"

	"github.com/goava/di"
	"github.com/labstack/echo/v4"
)

// appModules stores all necessary modules here for DI (Dependency Injection),
// then registered at `New` and `NewForTransaction`
var appModules = di.Options(
	di.Provide(customvalidator.NewValidator),
	di.Provide(databases.NewPostgresDB),
	di.Provide(echo.New),
)
