package services

import (
	"context"

	"github.com/irhamsahbana/gin-mongo/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	usercollection *mongo.Collection
	ctx context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserServiceContract {
	return &UserService {
		usercollection: usercollection,
		ctx: ctx,
	}
}

func (u *UserService) CreateUser(user *models.User) error {
	return nil
}

func (u *UserService) GetUser(name *string) (*models.User, error) {
	return nil, nil
}

func (u *UserService) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (u *UserService) UpdateUser(*models.User) error {
	return nil
}

func (u *UserService) DeleteUser(name *string) error {
	return nil
}