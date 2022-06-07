package services

import "github.com/xvbnm48/gin-crash-course/models"

type UserService interface {
	CreateUser(user *models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
