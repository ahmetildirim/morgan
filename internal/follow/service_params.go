package follow

import "github.com/google/uuid"

type CreateFollowServiceParams struct {
	FollowerID uuid.UUID
	FolloweeID uuid.UUID
}
