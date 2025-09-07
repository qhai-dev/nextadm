package service

import (
	"context"

	"github.com/nextadm/internal/user/domain/user/entity"
)

type UserService struct {
	//
}

func NewUserService() IUserService {
	return &UserService{
		//
	}
}

func (service *UserService) Create(ctx context.Context, req *CreateUserRequest) (user *entity.User, err error) {
	return &entity.User{}, nil
}
