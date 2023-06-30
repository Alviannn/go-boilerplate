package main

import (
	"fmt"
	"go-boilerplate/internal/domains/users"
	"go-boilerplate/internal/middlewares"
	"go-boilerplate/pkg/customvalidator"
	"go-boilerplate/pkg/databases"
	"log"
	"os"

	"github.com/goava/di"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProvideDIContainer() (container *di.Container, err error) {
	err = godotenv.Load()
	if err != nil {
		return
	}

	di.SetTracer(&di.StdTracer{})

	container, err = di.New(
		di.Provide(customvalidator.NewValidator),
		di.Provide(databases.NewPostgresDB),
		di.Provide(echo.New),

		// Include domains module
		users.Module,
	)
	return
}

// Insert all `SetupRouter` functions here.
var SetupRouterList = []di.Invocation{
	users.SetupRouter,
}

func main() {
	container, err := ProvideDIContainer()
	if err != nil {
		log.Fatal(err)
	}

	err = container.Invoke(func(e *echo.Echo) (err error) {
		// Force DB to load and test the connection.
		var gorm *gorm.DB
		if err = container.Resolve(&gorm); err != nil {
			return
		}

		// Override error handler middleware
		e.HTTPErrorHandler = middlewares.ErrorHandler()

		// Register all routes
		for _, setupRouterFunc := range SetupRouterList {
			if err = container.Invoke(setupRouterFunc); err != nil {
				return
			}
		}

		err = e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
		return
	})
	if err != nil {
		log.Fatal(err)
	}
}
