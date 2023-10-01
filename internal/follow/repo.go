package follow

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var (
	ErrNotFound = errors.New("follow not found")
)

type repo struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) *repo {
	return &repo{conn: conn}
}

func (r *repo) Create(ctx context.Context, follow *Follow) error {
	_, err := r.conn.Exec(ctx, "INSERT INTO follows (id, follower_id, followee_id, created_at) VALUES ($1, $2, $3, $4)",
		follow.ID, follow.FollowerID, follow.FolloweeID, follow.CreatedAt)

	return err
}

func (r *repo) FindByFollowerAndFollowee(ctx context.Context, followerID uuid.UUID, followeeID uuid.UUID) (*Follow, error) {
	follow := &Follow{}
	err := r.conn.QueryRow(ctx, "SELECT id, follower_id, followee_id, created_at FROM follows WHERE follower_id = $1 AND followee_id = $2", followerID, followeeID).
		Scan(&follow.ID, &follow.FollowerID, &follow.FolloweeID, &follow.CreatedAt)

	if err == pgx.ErrNoRows {
		return nil, ErrNotFound
	}

	return follow, err
}

func (r *repo) FindByFollower(ctx context.Context, followerID uuid.UUID) ([]*Follow, error) {
	rows, err := r.conn.Query(ctx, "SELECT id, follower_id, followee_id, created_at FROM follows WHERE follower_id = $1", followerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var follows []*Follow
	for rows.Next() {
		follow := &Follow{}
		err := rows.Scan(&follow.ID, &follow.FollowerID, &follow.FolloweeID, &follow.CreatedAt)
		if err != nil {
			return nil, err
		}

		follows = append(follows, follow)
	}

	return follows, nil
}
