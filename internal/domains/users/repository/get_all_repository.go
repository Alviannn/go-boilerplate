package users_repository

import (
	"go-boilerplate/internal/dtos"
	"go-boilerplate/internal/models"
)

func (r *RepositoryImpl) GetAllUsers(params dtos.GetAllUsersReq) (userList []models.User, err error) {
	query := r.DB

	if params.Email != "" {
		query = query.Where("email = ?", params.Email)
	}
	if params.FullName != "" {
		query = query.Where("full_name = ?", params.FullName)
	}
	if params.Username != "" {
		query = query.Where("username = ?", params.Username)
	}

	if params.Limit != 0 {
		query = query.Limit(int(params.Limit))
	}
	if params.Offset != 0 {
		query = query.Offset(params.Offset)
	}

	err = query.Find(&userList).Error
	return
}
