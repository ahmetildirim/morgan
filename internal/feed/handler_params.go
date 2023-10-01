package feed

import (
	"time"

	"github.com/google/uuid"
)

type GetFeedHandlerResponse struct {
	Posts []*GetFeedHandlerResponsePost `json:"posts"`
}

type GetFeedHandlerResponsePost struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
