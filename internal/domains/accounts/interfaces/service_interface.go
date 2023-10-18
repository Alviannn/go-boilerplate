package accounts_interfaces

import (
	"context"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
)

type Service interface {
	GetByID(ctx context.Context, params dtos.GetAccountReq) (account models.Account, err error)
	GetAll(ctx context.Context, params dtos.GetAllAccountsReq) (accounts []models.Account, err error)
	Register(ctx context.Context, params dtos.RegisterAccountReq) (err error)
}
