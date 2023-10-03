package post

import (
	"context"

	"github.com/google/uuid"
)

type repository interface {
	CreatePost(ctx context.Context, post *Post) error
	GetPostsByUserIDs(ctx context.Context, userIDs []uuid.UUID) ([]*Post, error)
	AddLike(ctx context.Context, postID uuid.UUID) error
}

type service interface {
	CreatePost(ctx context.Context, params *CreatePostServiceParams) (*Post, error)
}
