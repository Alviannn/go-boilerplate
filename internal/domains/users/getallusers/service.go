package getallusers

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
	"go-boilerplate/pkg/responses"
	"net/http"
)

type serviceImpl struct {
	Repository Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{Repository: repo}
}

func (s *serviceImpl) GetAllUsers(params dtos.GetAllUsersReq) (userList []models.User, err error) {
	userList, err = s.Repository.GetAllUsers(params)
	if err != nil {
		err = responses.NewError().
			WithSourceError(err).
			WithMessage("Failed to get all users.").
			WithCode(http.StatusInternalServerError)
	}
	return
}
