package auth

import (
	"context"

	"morgan.io/internal/user"
)

type service interface {
	Login(ctx context.Context, params *LoginServiceParams) (string, error)
}

type userService interface {
	Authenticate(ctx context.Context, email string, password string) (*user.User, error)
}
