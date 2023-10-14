package users_service

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

func (s *ServiceImpl) GetUser(ctx context.Context, params dtos.GetUserReq) (user models.User, err error) {
	user, err = s.Repository.GetUser(ctx, params.ID)
	if err != nil {
		newErr := responses.NewError().
			WithSourceError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage(fmt.Sprintf("Cannot fetch user with ID %d", params.ID))

		if errors.Is(err, gorm.ErrRecordNotFound) {
			newErr.
				WithCode(http.StatusNotFound).
				WithMessage(fmt.Sprintf("Cannot find user with ID %d", params.ID))
		}

		err = newErr
	}
	return
}
