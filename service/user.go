package service

import (
	config "back-end-website/config"
	"back-end-website/graph/model"
	"back-end-website/graph/mypkg"
	"context"
)

type IUserService interface {
	GetUser(ctx context.Context, id mypkg.UUID) (*model.User, error)
	GetUsers(ctx context.Context, limit int, offset int) ([]*model.User, error)
	UpdateUser(ctx context.Context, input *model.UpdateUserInput) (*model.User, error)
	DeleteUser(ctx context.Context, id mypkg.UUID) (*model.User, error)
	UpdateRole(ctx context.Context, role model.UserType, id mypkg.UUID) (*model.User, error)
}

type UserService struct {
	server *config.Server
}

func NewUserService(server *config.Server) *UserService {
	return &UserService{
		server: server,
	}
}
