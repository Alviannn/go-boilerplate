package accounts_service

import (
	"context"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
	"go-boilerplate/pkg/responses"
	"net/http"
)

func (s *serviceImpl) GetAll(ctx context.Context, params dtos.GetAllAccountsReq) (accounts []models.Account, err error) {
	accounts, err = s.PostgresRepository.GetAll(ctx, params)
	if err != nil {
		err = responses.NewError().
			WithSourceError(err).
			WithMessage("Failed to get all accounts.").
			WithCode(http.StatusInternalServerError)
	}
	return
}
