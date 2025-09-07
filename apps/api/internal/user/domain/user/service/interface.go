package service

import (
	"context"

	"github.com/nextadm/internal/user/domain/user/entity"
)

type IUserService interface {
	Create(ctx context.Context, req *CreateUserRequest) (user *entity.User, err error)
}

type CreateUserRequest struct {
	//
}
