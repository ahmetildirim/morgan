package feed

import (
	"context"

	"github.com/google/uuid"
	"morgan.io/internal/post"
)

type followService interface {
	GetFollowees(ctx context.Context, userID uuid.UUID) ([]uuid.UUID, error)
}

type postService interface {
	GetPostsByUserIDs(ctx context.Context, userIDs []uuid.UUID) ([]*post.Post, error)
}

type Service struct {
	followSvc followService
	postSvc   postService
}

func NewService(followSvc followService, postSvc postService) *Service {
	return &Service{
		followSvc: followSvc,
		postSvc:   postSvc,
	}
}

func (s *Service) GetFeed(ctx context.Context, userID uuid.UUID) (*Feed, error) {
	followees, err := s.followSvc.GetFollowees(ctx, userID)
	if err != nil {
		return nil, err
	}

	if len(followees) == 0 {
		return NewFeed(nil), nil
	}

	posts, err := s.postSvc.GetPostsByUserIDs(ctx, followees)
	if err != nil {
		return nil, err
	}

	feed := NewFeed(posts)

	return feed, nil
}
