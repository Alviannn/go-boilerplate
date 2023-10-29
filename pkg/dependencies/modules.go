package dependencies

import (
	"go-boilerplate/pkg/customvalidator"
	"go-boilerplate/pkg/databases"

	"github.com/defval/di"
)

// appModules stores all necessary modules here for DI (Dependency Injection),
// then registered at `New` and `NewForTransaction`
var appModules = di.Options(
	di.Provide(customvalidator.New),
	di.Provide(databases.NewMySQLDB),
)
