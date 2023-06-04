package registeruser

import (
	"errors"
	"go-boilerplate/internal/constants"
	"go-boilerplate/internal/dtos"

	"golang.org/x/crypto/bcrypt"
)

type serviceImpl struct {
	Repository Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{Repository: repo}
}

func (s *serviceImpl) RegisterUser(params dtos.RegisterUserReq) (err error) {
	if s.Repository.IsUserExistByEmail(params.Email) {
		return errors.New("email is already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(params.Password), constants.DefaultHashCost)
	if err != nil {
		return
	}

	params.Password = string(hashed)
	return s.Repository.RegisterUser(params)
}
