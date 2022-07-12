package services

import "github.com/irhamsahbana/gin-mongo/models"

type UserServiceContract interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}