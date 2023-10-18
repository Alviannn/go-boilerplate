package accounts_service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
	"go-boilerplate/pkg/responses"

	"gorm.io/gorm"
)

func (s *ServiceImpl) GetByID(ctx context.Context, params dtos.GetAccountReq) (account models.Account, err error) {
	account, err = s.Repository.GetByID(ctx, params.ID)
	if err != nil {
		newErr := responses.NewError().
			WithSourceError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage(fmt.Sprintf("Failed to get account with ID %d", params.ID))

		if errors.Is(err, gorm.ErrRecordNotFound) {
			newErr.
				WithCode(http.StatusNotFound).
				WithMessage(fmt.Sprintf("Cannot find account with ID %d", params.ID))
		}

		err = newErr
	}
	return
}
