package domain

import "context"

type User struct {
	ID    uint   `json:"id" gorm:"unique;not null"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository interface {
	Save(ctx context.Context, user User) (User, error)
}

type UserService interface {
	Save(ctx context.Context, user User) (User, error)
}
