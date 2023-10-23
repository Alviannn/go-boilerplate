package accounts_interfaces

import (
	"context"
	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
)

type Service interface {
	GetByID(ctx context.Context, params dtos.GetAccountReq) (account models_mysql.Account, err error)
	GetAll(ctx context.Context, params dtos.GetAllAccountsReq) (accounts []models_mysql.Account, err error)
	Register(ctx context.Context, params dtos.RegisterAccountReq) (err error)
}
