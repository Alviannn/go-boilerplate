package users

import (
	"go-boilerplate/internal/domains/users/getallusers"
	"go-boilerplate/internal/domains/users/getuser"
	"go-boilerplate/internal/domains/users/registeruser"

	"github.com/goava/di"
	"github.com/labstack/echo/v4"
)

type UsersRouterParams struct {
	di.Inject

	Echo *echo.Echo

	GetUser      getuser.Handler
	GetAllUsers  getallusers.Handler
	RegisterUser registeruser.Handler
}

func SetupRouter(r UsersRouterParams) {
	router := r.Echo.Group("users")

	router.GET("/:id", r.GetUser.Handle)
	router.GET("", r.GetAllUsers.Handle)
	router.POST("", r.RegisterUser.Handle)
}
