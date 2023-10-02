package comment

import "context"

type repository interface {
	Create(ctx context.Context, comment *Comment) error
}

type service interface {
	CreateComment(ctx context.Context, params *CreateCommentServiceParams) (*Comment, error)
}
