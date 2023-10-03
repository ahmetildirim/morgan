package like

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrPostNotFound = errors.New("post not found")
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

	return s.repo.Create(ctx, like)
}

func (s *Service) FindByPostID(ctx context.Context, postID uuid.UUID) ([]*Like, error) {
	return s.repo.FindByPostID(ctx, postID)
}
