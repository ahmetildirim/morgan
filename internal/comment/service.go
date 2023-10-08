package comment

import (
	"context"

	"github.com/google/uuid"
)

type CreateCommentServiceParams struct {
	PostID  uuid.UUID
	OwnerID uuid.UUID
	Content string
}

type repository interface {
	Create(ctx context.Context, comment *Comment) error
}

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateComment(ctx context.Context, params *CreateCommentServiceParams) (*Comment, error) {
	comment, err := NewComment(params.PostID, params.OwnerID, params.Content)
	if err != nil {
		return nil, err
	}

	err = s.repo.Create(ctx, comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}
