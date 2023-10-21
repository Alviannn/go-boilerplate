package accounts_service

import (
	"context"
	"fmt"
	"net/http"

	"go-boilerplate/internal/dtos"
	postgres_models "go-boilerplate/internal/models/postgres"
	"go-boilerplate/pkg/responses"
)

func (s *serviceImpl) GetByID(ctx context.Context, params dtos.GetAccountReq) (account postgres_models.Account, err error) {
	account, err = s.PostgresRepository.GetByID(ctx, params.ID)
	if err != nil {
		err = responses.NewError().
			WithSourceError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage(fmt.Sprintf("Failed to get account with ID %d", params.ID))
		return
	}
	if !account.IsExist() {
		err = responses.NewError().
			WithCode(http.StatusNotFound).
			WithMessage(fmt.Sprintf("Cannot find account with ID %d", params.ID))
	}

	return
}
