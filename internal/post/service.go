package post

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePost(ctx context.Context, params *CreatePostServiceParams) (*Post, error) {
	post, err := NewPost(params.OwnerID, params.Content)
	if err != nil {
		return nil, err
	}

	err = s.repo.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *Service) GetPostsByUserIDs(ctx context.Context, userIDs []uuid.UUID) ([]*Post, error) {
	posts, err := s.repo.GetPostsByUserIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *Service) AddLike(ctx context.Context, postID uuid.UUID) error {
	err := s.repo.AddLike(ctx, postID)
	if err != nil {
		return err
	}

	return nil
}
