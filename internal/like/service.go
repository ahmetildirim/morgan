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
	repo    repository
	postSvc postService
}

func NewService(repo repository, postSvc postService) *Service {
	return &Service{
		repo:    repo,
		postSvc: postSvc,
	}
}

func (s *Service) Create(ctx context.Context, postID, ownerID uuid.UUID) error {
	like, err := NewLike(postID, ownerID)
	if err != nil {
		return err
	}

	err = s.repo.Create(ctx, like)
	if err != nil {
		return err
	}

	err = s.postSvc.AddLike(ctx, postID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) FindByPostID(ctx context.Context, postID uuid.UUID) ([]*Like, error) {
	return s.repo.FindByPostID(ctx, postID)
}
