package users_interfaces

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
)

type Service interface {
	GetUser(params dtos.GetUserReq) (user models.User, err error)
	GetAllUsers(params dtos.GetAllUsersReq) (userList []models.User, err error)
	RegisterUser(params dtos.RegisterUserReq) (err error)
}
