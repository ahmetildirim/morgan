package post

import (
	"context"
	"errors"

	"github.com/google/uuid"
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

func (r *repo) GetPostsByUserIDs(ctx context.Context, userIDs []uuid.UUID) ([]*Post, error) {
	rows, err := r.conn.Query(ctx, "SELECT id, owner_id, content, created_at, updated_at FROM posts WHERE owner_id = ANY($1)", userIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*Post
	for rows.Next() {
		post := &Post{}
		err := rows.Scan(&post.ID, &post.OwnerID, &post.Content, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (r *repo) AddLike(ctx context.Context, postID uuid.UUID) error {
	_, err := r.conn.Exec(ctx, "UPDATE posts SET likes = likes + 1 WHERE id = $1", postID)

	return err
}
