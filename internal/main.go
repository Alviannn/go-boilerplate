package main

import (
	"fmt"
	"go-boilerplate/internal/pkg/databases"
	"log"
	"os"

	"github.com/goava/di"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func ProvideDIContainer() (container *di.Container, err error) {
	err = godotenv.Load()
	if err != nil {
		return
	}

	container, err = di.New(
		di.Provide(databases.NewPostgresDB),
		di.Provide(echo.New),
	)
	return
}

func main() {
	di.SetTracer(&di.StdTracer{})

	container, err := ProvideDIContainer()
	if err != nil {
		log.Fatal(err)
	}

	container.Invoke(func(e *echo.Echo) {
		e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
	})
}
