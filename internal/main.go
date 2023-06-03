package main

import (
	"fmt"
	"go-boilerplate/internal/helpers"
	"go-boilerplate/pkg/databases"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

func ProvideDIContainer() (container *dig.Container, err error) {
	err = godotenv.Load()
	if err != nil {
		return
	}

	container = dig.New()
	err = helpers.MultiProvideDI(container, []any{
		databases.NewPostgresDB,
		echo.New,
	})

	return
}

func main() {
	container, err := ProvideDIContainer()
	if err != nil {
		log.Fatal(err)
	}

	container.Invoke(func(e *echo.Echo) {
		e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
	})
}
