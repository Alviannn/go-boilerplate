package getuser

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
)

type serviceImpl struct {
	Repository Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{Repository: repo}
}

func (s *serviceImpl) GetUser(params dtos.GetUserReq) (user models.User, err error) {
	return s.Repository.GetUser(params.UserID)
}
