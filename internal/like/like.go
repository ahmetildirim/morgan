package like

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrPostEmpty  = errors.New("post id is empty")
	ErrOwnerEmpty = errors.New("owner id is empty")
)

type Like struct {
	ID        uuid.UUID
	PostID    uuid.UUID
	OwnerID   uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewLike(postID, ownerID uuid.UUID) (*Like, error) {
	if postID == uuid.Nil {
		return nil, ErrPostEmpty
	}

	if ownerID == uuid.Nil {
		return nil, ErrOwnerEmpty
	}

	return &Like{
		ID:        uuid.New(),
		PostID:    postID,
		OwnerID:   ownerID,
		CreatedAt: time.Now(),
	}, nil
}
