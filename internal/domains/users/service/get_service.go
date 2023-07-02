package users_service

import (
	"errors"
	"fmt"
	"net/http"

	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
	"go-boilerplate/pkg/responses"

	"gorm.io/gorm"
)

func (s *ServiceImpl) GetUser(params dtos.GetUserReq) (user models.User, err error) {
	user, err = s.Repository.GetUser(params.UserID)
	if err != nil {
		newErr := responses.NewError().
			WithSourceError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage(fmt.Sprintf("Cannot fetch user with ID %s", params.UserID))

		if errors.Is(err, gorm.ErrRecordNotFound) {
			newErr.
				WithCode(http.StatusNotFound).
				WithMessage(fmt.Sprintf("Cannot find user with ID %s", params.UserID))
		}

		err = newErr
	}
	return
}
