package user_service

import (
	"gofinn/internal/models"
)

type UserService interface {
	ListUsers() ([]models.User, error)
	CreateUser(user models.UserFields) (models.User, error)
	GetUser(userID int64) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(userID int64) error
}
