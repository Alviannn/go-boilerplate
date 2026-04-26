package controllers_rest

import (
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
)

func Register(injector do.Injector, app *echo.Echo) error {
	constructorList := []do.Provider[Controller]{
		NewHealth,
		NewAccounts,
	}

	for _, constructor := range constructorList {
		controller, err := constructor(injector)
		if err != nil {
			return err
		}
		controller.SetupRouter(app)
	}
	return nil
}
