package users

import (
	"go-boilerplate/internal/domains/users/getuser"

	"github.com/goava/di"
	"github.com/labstack/echo/v4"
)

type UsersRouterParams struct {
	di.Inject

	Echo    *echo.Echo
	GetUser getuser.Handler
}

func SetupRouter(r UsersRouterParams) {
	router := r.Echo.Group("users")

	router.GET("/:id", r.GetUser.Handle)
}
