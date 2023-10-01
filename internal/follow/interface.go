package follow

import (
	"context"

	"github.com/google/uuid"
	"morgan.io/internal/user"
)

type repository interface {
	Create(ctx context.Context, follow *Follow) error
	FindByFollowerAndFollowee(ctx context.Context, followerID uuid.UUID, followeeID uuid.UUID) (*Follow, error)
}

type service interface {
	Follow(ctx context.Context, params *CreateFollowServiceParams) error
}

type userService interface {
	GetUser(ctx context.Context, id uuid.UUID) (*user.User, error)
}
