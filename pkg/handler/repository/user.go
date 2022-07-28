package repository

import (
	"context"
	"go-wire-architecture/pkg/model"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) model.UserRepository {
	return &userDatabase{DB}
}

func (c *userDatabase) Save(ctx context.Context, user model.Users) (model.Users, error) {
	err := c.DB.Save(&user).Error

	return user, err
}
