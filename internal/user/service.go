package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
)

type CreateServiceParams struct {
	Email    string
	Password string
}

type AuthenticateServiceParams struct {
	Email    string
	Password string
}

type repository interface {
	CreateUser(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
}

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUser(ctx context.Context, params *CreateServiceParams) (*User, error) {
	_, err := s.repo.FindByEmail(ctx, params.Email)
	if err != ErrNotFound {
		return nil, ErrEmailAlreadyExists
	}

	user, err := NewUser(params.Email, params.Password)
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Authenticate(ctx context.Context, email string, password string) (*User, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !user.CheckPassword(password) {
		return nil, ErrInvalidPassword
	}

	return user, nil
}

func (s *Service) GetUser(ctx context.Context, id uuid.UUID) (*User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
