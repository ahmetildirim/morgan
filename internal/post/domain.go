package post

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNoContent = errors.New("content cannot be empty")
	ErrNoOwnerID = errors.New("ownerID cannot be empty")
)

type Post struct {
	ID        uuid.UUID
	OwnerID   uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPost(ownerID uuid.UUID, content string) (*Post, error) {
	if content == "" {
		return nil, ErrNoContent
	}

	if ownerID == uuid.Nil {
		return nil, ErrNoOwnerID
	}

	return &Post{
		ID:      uuid.New(),
		OwnerID: ownerID,
		Content: content,
	}, nil
}
