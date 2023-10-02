package comment

import (
	"time"

	"github.com/google/uuid"
)

type CreateCommentHandlerParams struct {
	PostID  uuid.UUID `json:"post_id"`
	Content string    `json:"content"`
}

type CreateCommentHandlerResponse struct {
	ID        uuid.UUID `json:"id"`
	PostID    uuid.UUID `json:"post_id"`
	OwnerID   uuid.UUID `json:"owner_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
