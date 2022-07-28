package model

import (
	"context"
)

type Users struct {
	ID    uint   `json:"id" gorm:"unique;not null"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository interface {
	Save(ctx context.Context, user Users) (Users, error)
}
