package application

import (
	userv1 "github.com/nextadm/gen/user/v1"
	"github.com/nextadm/internal/user/domain/user/service"
)

type UserApplication struct {
	userv1.UnimplementedUserServiceServer
	userService service.UserService
}

func NewUserApplication() *UserApplication {
	// TODO
	return &UserApplication{}
}