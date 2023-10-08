package like

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrLikeAlreadyExists = errors.New("like already exists")
)

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, postID, ownerID uuid.UUID) error {
	like, err := NewLike(postID, ownerID)
	if err != nil {
		return err
	}

	exists, err := s.repo.Exists(ctx, postID, ownerID)
	if err != nil {
		return err
	}

	if exists {
		return ErrLikeAlreadyExists
	}

	err = s.repo.Create(ctx, like)
	if err != nil {
		return err
	}

	return nil
}
