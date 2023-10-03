package like

import (
	"context"

	"github.com/google/uuid"
)

type service interface {
	Create(ctx context.Context, postID, ownerID uuid.UUID) error
	FindByPostID(ctx context.Context, postID uuid.UUID) ([]*Like, error)
}

type postService interface {
	AddLike(ctx context.Context, postID uuid.UUID) error
}

type repository interface {
	Create(ctx context.Context, like *Like) error
	FindByPostID(ctx context.Context, postID uuid.UUID) ([]*Like, error)
}
