package user

import (
    "context"
)

type Repository interface {
    Save(ctx context.Context, user Users) (Users, error)
}

type Service interface {
    Save(ctx context.Context, user Users) (Users, error)
}
