package getuser

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
)

type Service interface {
	GetUser(params dtos.GetUserReq) (user models.User, err error)
}

type Repository interface {
	GetUser(userID string) (user models.User, err error)
}
