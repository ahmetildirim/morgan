package auth

import (
	"context"

	"morgan.io/internal/user"
)

type LoginServiceParams struct {
	Email    string
	Password string
}

type userService interface {
	Authenticate(ctx context.Context, email string, password string) (*user.User, error)
}

type Service struct {
	userService userService
	secretKey   string
}

func NewService(userService userService, secretKey string) *Service {
	return &Service{
		userService: userService,
		secretKey:   secretKey,
	}
}

func (s Service) Login(ctx context.Context, params *LoginServiceParams) (*Token, error) {
	user, err := s.userService.Authenticate(ctx, params.Email, params.Password)
	if err != nil {
		return nil, err
	}

	token, err := NewToken(user, s.secretKey)
	if err != nil {
		return nil, err
	}

	return token, nil
}
