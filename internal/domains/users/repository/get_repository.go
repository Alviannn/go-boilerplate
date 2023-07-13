package users_repository

import "go-boilerplate/internal/models"

func (r *RepositoryImpl) GetUser(userID int64) (users models.User, err error) {
	err = r.DB.First(&users, "id = ?", userID).Error
	return
}
