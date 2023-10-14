package users_service

import (
	"context"
	"go-boilerplate/internal/constants"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/responses"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (s *ServiceImpl) RegisterUser(ctx context.Context, params dtos.RegisterUserReq) (err error) {
	if s.Repository.IsUserExistByEmail(ctx, params.Email) {
		err = responses.NewError().
			WithCode(http.StatusBadRequest).
			WithMessage("User is already registered.")
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(params.Password), constants.DefaultHashCost)
	if err != nil {
		err = responses.NewError().
			WithSourceError(err).
			WithMessage("Failed to hash password.").
			WithCode(http.StatusBadRequest)
		return
	}

	params.Password = string(hashed)
	if err = s.Repository.RegisterUser(ctx, params); err != nil {
		err = responses.NewError().
			WithSourceError(err).
			WithMessage("Failed to register new user.").
			WithCode(http.StatusBadRequest)
	}

	return
}
