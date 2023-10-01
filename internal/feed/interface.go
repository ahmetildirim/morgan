package feed

import (
	"context"

	"github.com/google/uuid"
	"morgan.io/internal/post"
)

type service interface {
	GetFeed(ctx context.Context, userID uuid.UUID) (*Feed, error)
}

type followService interface {
	GetFollowees(ctx context.Context, userID uuid.UUID) ([]uuid.UUID, error)
}

type postService interface {
	GetPostsByUserIDs(ctx context.Context, userIDs []uuid.UUID) ([]*post.Post, error)
}
