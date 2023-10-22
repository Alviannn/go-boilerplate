package accounts_service

import (
	"context"
	"go-boilerplate/internal/constants"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/customerror"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (s *serviceImpl) Register(ctx context.Context, params dtos.RegisterAccountReq) (err error) {
	if s.MySQLRepository.IsExistByEmail(ctx, params.Email) {
		err = customerror.New().
			WithCode(http.StatusBadRequest).
			WithMessage("Account is already registered.")
		return
	}

	params.Password, err = s.hashPassword(params.Password)
	if err != nil {
		return
	}

	if err = s.MySQLRepository.Register(ctx, params); err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithMessage("Failed to register new account.").
			WithCode(http.StatusBadRequest)
	}
	return
}

func (s *serviceImpl) hashPassword(password string) (hashed string, err error) {
	rawHashed, err := bcrypt.GenerateFromPassword([]byte(password), constants.DefaultHashCost)
	if err != nil {
		err = customerror.New().
			WithSourceError(err).
			WithMessage("Failed to hash password.").
			WithCode(http.StatusBadRequest)
		return
	}

	hashed = string(rawHashed)
	return
}
