package services

import (
	"context"
	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
)

type AccountsService interface {
	GetByID(ctx context.Context, params dtos.AccountGetRequest) (account models_mysql.Account, err error)
	GetAll(ctx context.Context, params dtos.AccountGetAllRequest) (accounts []models_mysql.Account, err error)
	Register(ctx context.Context, params dtos.AccountRegisterRequest) (err error)
}
