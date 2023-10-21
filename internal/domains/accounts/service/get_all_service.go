package accounts_service

import (
	"context"
	"go-boilerplate/internal/dtos"
	postgres_models "go-boilerplate/internal/models/postgres"
	"go-boilerplate/pkg/responses"
	"net/http"
)

func (s *serviceImpl) GetAll(ctx context.Context, params dtos.GetAllAccountsReq) (accounts []postgres_models.Account, err error) {
	accounts, err = s.PostgresRepository.GetAll(ctx, params)
	if err != nil {
		err = responses.NewError().
			WithSourceError(err).
			WithMessage("Failed to get all accounts.").
			WithCode(http.StatusInternalServerError)
	}
	return
}
