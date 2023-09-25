package user

import "context"

type repository interface {
	CreateUser(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
}

type service interface {
	CreateUser(ctx context.Context, params *CreateServiceParams) (*User, error)
	Authenticate(ctx context.Context, email string, password string) (*User, error)
}
