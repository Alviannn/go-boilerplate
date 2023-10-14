package users_interfaces

import (
	"context"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
)

type Repository interface {
	GetUser(ctx context.Context, userID int64) (user models.User, err error)
	GetAllUsers(ctx context.Context, params dtos.GetAllUsersReq) (userList []models.User, err error)
	RegisterUser(ctx context.Context, params dtos.RegisterUserReq) error
	IsUserExistByEmail(ctx context.Context, email string) bool
}
