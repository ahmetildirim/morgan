package post

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

var (
	ErrNotFound = errors.New("user not found")
)

type repo struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) *repo {
	return &repo{conn: conn}
}

func (r *repo) CreatePost(ctx context.Context, post *Post) error {
	_, err := r.conn.Exec(ctx, "INSERT INTO posts (id, owner_id, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		post.ID, post.OwnerID, post.Content, post.CreatedAt, post.UpdatedAt)

	return err
}
