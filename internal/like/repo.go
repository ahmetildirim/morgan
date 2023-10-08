package like

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type repo struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) *repo {
	return &repo{conn: conn}
}

func (r *repo) Create(ctx context.Context, like *Like) error {
	_, err := r.conn.Exec(ctx, "INSERT INTO likes (id, post_id, owner_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		like.ID, like.PostID, like.OwnerID, like.CreatedAt, like.UpdatedAt)

	return err
}

func (r *repo) Exists(ctx context.Context, postID, ownerID uuid.UUID) (bool, error) {
	var exists bool
	err := r.conn.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM likes WHERE post_id = $1 AND owner_id = $2)", postID, ownerID).Scan(&exists)

	return exists, err
}

func (r *repo) FindByPostID(ctx context.Context, postID uuid.UUID) ([]*Like, error) {
	rows, err := r.conn.Query(ctx, "SELECT id, post_id, owner_id, created_at, updated_at FROM likes WHERE post_id = $1", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likes []*Like
	for rows.Next() {
		like := &Like{}
		err := rows.Scan(&like.ID, &like.PostID, &like.OwnerID, &like.CreatedAt, &like.UpdatedAt)
		if err != nil {
			return nil, err
		}

		likes = append(likes, like)
	}

	return likes, nil
}
