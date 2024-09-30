package services

import (
	"context"
	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
)

type AccountsService interface {
	GetByID(ctx context.Context, params dtos.AccountGetReq) (account models_mysql.Account, err error)
	GetAll(ctx context.Context, params dtos.AccountGetAllReq) (accounts []models_mysql.Account, err error)
	Register(ctx context.Context, params dtos.AccountRegisterReq) (err error)
}
