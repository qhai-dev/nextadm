package repo

import (
	"context"

	"github.com/nextadm/internal/user/domain/user/entity"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *entity.User) (userID int, err error)
}

// TODO
