package comment

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEmptyContent = errors.New("content is empty")
	ErrPostEmpty    = errors.New("post id is empty")
	ErrOwnerEmpty   = errors.New("owner id is empty")
)

type Comment struct {
	ID        uuid.UUID
	PostID    uuid.UUID
	OwnerID   uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdateAt  time.Time
}

func NewComment(postID, ownerID uuid.UUID, content string) (*Comment, error) {
	if content == "" {
		return nil, ErrEmptyContent
	}

	if postID == uuid.Nil {
		return nil, ErrPostEmpty
	}

	if ownerID == uuid.Nil {
		return nil, ErrOwnerEmpty
	}

	return &Comment{
		ID:        uuid.New(),
		PostID:    postID,
		OwnerID:   ownerID,
		Content:   content,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}, nil
}
