package services

import (
	"context"

	"github.com/xvbnm48/gin-crash-course/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserSerficeImpl{
		userCollection: usercollection,
		ctx:            ctx,
	}
}

type UserSerficeImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func (u *UserSerficeImpl) CreateUser(user *models.User) error {
	return nil
}

func (u *UserSerficeImpl) GetUser(name *string) (*models.User, error) {
	return nil, nil
}

func (u *UserSerficeImpl) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (u *UserSerficeImpl) UpdateUser(user *models.User) error {
	return nil
}

func (u *UserSerficeImpl) DeleteUser(name *string) error {
	return nil
}
