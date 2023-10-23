package accounts_service

import (
	"context"
	"fmt"
	"net/http"

	"go-boilerplate/internal/dtos"
	models_mysql "go-boilerplate/internal/models/mysql"
	"go-boilerplate/pkg/customerror"
)

func (s *serviceImpl) GetByID(ctx context.Context, params dtos.GetAccountReq) (account models_mysql.Account, err error) {
	account, err = s.RepositoryMySQL.GetByID(ctx, params.ID)
	if err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage(fmt.Sprintf("Failed to get account with ID %d", params.ID))
		return
	}
	if !account.IsExist() {
		err = customerror.New().
			WithCode(http.StatusNotFound).
			WithMessage(fmt.Sprintf("Cannot find account with ID %d", params.ID))
	}

	return
}
