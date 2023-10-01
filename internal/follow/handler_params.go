package follow

import "github.com/google/uuid"

type CreateFollowHandlerParams struct {
	FolloweeID uuid.UUID `json:"followee_id"`
}
