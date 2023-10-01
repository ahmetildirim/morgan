package post

import "context"

type repository interface {
	CreatePost(ctx context.Context, post *Post) error
}

type service interface {
	CreatePost(ctx context.Context, params *CreatePostServiceParams) (*Post, error)
}
