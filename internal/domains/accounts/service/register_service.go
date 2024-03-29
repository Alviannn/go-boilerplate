package accounts_service

import (
	"context"
	"go-boilerplate/internal/constants"
	"go-boilerplate/internal/dtos"
	"go-boilerplate/pkg/customerror"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (s *serviceImpl) Register(ctx context.Context, params dtos.AccountRegisterRequest) (err error) {
	isExist, err := s.RepositoryMySQL.ExistByEmail(ctx, params.Email)
	if err != nil {
		return
	}
	if isExist {
		err = customerror.New().
			WithCode(http.StatusConflict).
			WithMessage("Account with this email already exists.")
		return
	}

	params.Password, err = s.hashPassword(params.Password)
	if err != nil {
		return
	}

	if err = s.RepositoryMySQL.Register(ctx, params); err != nil {
		return
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
