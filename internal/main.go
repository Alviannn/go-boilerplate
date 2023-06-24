package main

import (
	"fmt"
	"go-boilerplate/internal/domains/users"
	"go-boilerplate/internal/middlewares"
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

	// Force DB to load and test the connection.
	var gorm *gorm.DB
	if err := container.Resolve(&gorm); err != nil {
		return
	}

	container.Invoke(func(e *echo.Echo) {
		e.HTTPErrorHandler = middlewares.ErrorHandler()

		// Register all routes
		for _, setupRouterFunc := range SetupRouterList {
			di.Invoke(setupRouterFunc)
		}

		e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
	})
}
