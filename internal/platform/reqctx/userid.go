package reqctx

import (
	"context"

	"github.com/google/uuid"
)

type userIDKey struct{}

// WithUserID returns a new context with the given user ID.
func WithUserID(ctx context.Context, userID uuid.UUID) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

// UserIDFromContext returns the user ID from the given context.
func UserIDFromContext(ctx context.Context) (uuid.UUID, bool) {
	userID := ctx.Value(userIDKey{})
	if userID == nil {
		return uuid.Nil, false
	}

	return userID.(uuid.UUID), true
}
