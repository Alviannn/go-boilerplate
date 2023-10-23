package accounts_service

import (
	"context"
	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
	"go-boilerplate/pkg/customerror"
	"net/http"
)

func (s *serviceImpl) GetAll(ctx context.Context, params dtos.GetAllAccountsReq) (accounts []models_mysql.Account, err error) {
	accounts, err = s.RepositoryMySQL.GetAll(ctx, params)
	if err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithMessage("Failed to get all accounts.").
			WithCode(http.StatusInternalServerError)
	}
	return
}
