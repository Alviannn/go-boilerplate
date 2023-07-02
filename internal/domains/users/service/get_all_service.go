package users_service

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
	"go-boilerplate/pkg/responses"
	"net/http"
)

func (s *ServiceImpl) GetAllUsers(params dtos.GetAllUsersReq) (userList []models.User, err error) {
	userList, err = s.Repository.GetAllUsers(params)
	if err != nil {
		err = responses.NewError().
			WithSourceError(err).
			WithMessage("Failed to get all users.").
			WithCode(http.StatusInternalServerError)
	}
	return
}
