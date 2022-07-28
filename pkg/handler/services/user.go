package services

import (
	"context"
	"go-wire-architecture/pkg/model"
)

type UserService interface {
	Save(ctx context.Context, user model.Users) (model.Users, error)
}
