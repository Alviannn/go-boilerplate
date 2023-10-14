package users_interfaces

import (
	"context"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
)

type Service interface {
	GetUser(ctx context.Context, params dtos.GetUserReq) (user models.User, err error)
	GetAllUsers(ctx context.Context, params dtos.GetAllUsersReq) (userList []models.User, err error)
	RegisterUser(ctx context.Context, params dtos.RegisterUserReq) (err error)
}
