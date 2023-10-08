package like

import (
	"context"

	"github.com/google/uuid"
)

type postService interface {
	AddLike(ctx context.Context, postID uuid.UUID) error
}

type repository interface {
	Create(ctx context.Context, like *Like) error
	Exists(ctx context.Context, postID, ownerID uuid.UUID) (bool, error)
}
