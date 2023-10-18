package accounts_interfaces

import (
	"context"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
)

type Repository interface {
	Get(ctx context.Context, accountID int64) (account models.Account, err error)
	GetAll(ctx context.Context, params dtos.GetAllAccountsReq) (accounts []models.Account, err error)
	Register(ctx context.Context, params dtos.RegisterAccountReq) error
	IsExistByEmail(ctx context.Context, email string) bool
}
