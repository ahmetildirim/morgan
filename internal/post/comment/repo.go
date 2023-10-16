package comment

import (
	"context"
	"errors"

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

func (r *repo) Create(ctx context.Context, comment *Comment) error {
	_, err := r.conn.Exec(ctx, "INSERT INTO comments (id, post_id, owner_id, content, created_at) VALUES ($1, $2, $3, $4, $5)",
		comment.ID, comment.PostID, comment.OwnerID, comment.Content, comment.CreatedAt)

	return err
}
