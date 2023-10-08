package follow

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNoUserID = errors.New("userID cannot be empty")
)

type Follow struct {
	ID         uuid.UUID
	FollowerID uuid.UUID
	FolloweeID uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewFollow(followerID uuid.UUID, followeeID uuid.UUID) (*Follow, error) {
	if followerID == uuid.Nil {
		return nil, ErrNoUserID
	}

	if followeeID == uuid.Nil {
		return nil, ErrNoUserID
	}

	return &Follow{
		ID:         uuid.New(),
		FollowerID: followerID,
		FolloweeID: followeeID,
		CreatedAt:  time.Now(),
	}, nil
}
