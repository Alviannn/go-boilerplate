package getallusers

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
)

type Service interface {
	GetAllUsers(params dtos.GetAllUsersReq) (userList []models.User, err error)
}

type Repository interface {
	GetAllUsers(params dtos.GetAllUsersReq) (userList []models.User, err error)
}
