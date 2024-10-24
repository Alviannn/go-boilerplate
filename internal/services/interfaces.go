package services

import (
	"context"
	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
)

type Accounts interface {
	GetByID(ctx context.Context, param dtos.AccountGetReq) (account models_mysql.Account, err error)
	GetAll(ctx context.Context, param dtos.AccountGetAllReq) (accounts []models_mysql.Account, err error)
	Register(ctx context.Context, param dtos.AccountRegisterReq) (err error)
}
