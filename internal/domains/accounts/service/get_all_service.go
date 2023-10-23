package accounts_service

import (
	"context"
	"go-boilerplate/internal/dtos"
	mysql_models "go-boilerplate/internal/models/mysql"
	"go-boilerplate/pkg/customerror"
	"net/http"
)

func (s *serviceImpl) GetAll(ctx context.Context, params dtos.GetAllAccountsReq) (accounts []mysql_models.Account, err error) {
	accounts, err = s.RepositoryMySQL.GetAll(ctx, params)
	if err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithMessage("Failed to get all accounts.").
			WithCode(http.StatusInternalServerError)
	}
	return
}
