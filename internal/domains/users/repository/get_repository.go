package users_repository

import (
	"context"
	"go-boilerplate/internal/models"
)

func (r *RepositoryImpl) GetUser(ctx context.Context, userID int64) (users models.User, err error) {
	err = r.DB.WithContext(ctx).First(&users, "id = ?", userID).Error
	return
}
