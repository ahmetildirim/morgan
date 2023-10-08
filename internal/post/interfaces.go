package post

import (
	"context"

	"github.com/google/uuid"
)

type repository interface {
	CreatePost(ctx context.Context, post *Post) error
	GetPostsByUserIDs(ctx context.Context, userIDs []uuid.UUID) ([]*Post, error)
	AddLike(ctx context.Context, postID uuid.UUID) error
	Exists(ctx context.Context, postID uuid.UUID) (bool, error)
}

type service interface {
	CreatePost(ctx context.Context, params *CreatePostServiceParams) (*Post, error)
	AddLike(ctx context.Context, postID, ownerID uuid.UUID) error
}

type likeService interface {
	Create(ctx context.Context, postID, ownerID uuid.UUID) error
}
