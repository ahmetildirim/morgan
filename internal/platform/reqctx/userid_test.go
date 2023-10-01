package reqctx_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"morgan.io/internal/platform/reqctx"
)

func TestUserID(t *testing.T) {
	ctx := context.Background()
	userID := uuid.New()

	ctx = reqctx.WithUserID(ctx, userID)

	result, ok := reqctx.UserIDFromContext(ctx)
	if !ok {
		t.Error("expected user ID in context")
	}

	if result != userID {
		t.Errorf("expected user ID %v, but got %v", userID, result)
	}
}
