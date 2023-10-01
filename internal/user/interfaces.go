package user

import (
	"context"

	"github.com/google/uuid"
)

type repository interface {
	CreateUser(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
}

type service interface {
	CreateUser(ctx context.Context, params *CreateServiceParams) (*User, error)
	Authenticate(ctx context.Context, email string, password string) (*User, error)
}
