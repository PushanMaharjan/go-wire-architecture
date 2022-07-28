package handler

import (
	"context"
	"go-wire-architecture/pkg/handler/services"
	"go-wire-architecture/pkg/model"
)

type userHandler struct {
	userRepo model.UserRepository
}

func NewUserHandler(repo model.UserRepository) services.UserService {
	return &userHandler{
		userRepo: repo,
	}
}

func (c *userHandler) Save(ctx context.Context, user model.Users) (model.Users, error) {
	user, err := c.userRepo.Save(ctx, user)

	return user, err
}
