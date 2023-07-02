package users_interfaces

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
)

type Repository interface {
	GetUser(userID string) (user models.User, err error)
	GetAllUsers(params dtos.GetAllUsersReq) (userList []models.User, err error)
	RegisterUser(params dtos.RegisterUserReq) error
	IsUserExistByEmail(email string) bool
}
