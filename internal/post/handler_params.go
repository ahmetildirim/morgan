package post

import "github.com/google/uuid"

type CreatePostHandlerParams struct {
	Content string `json:"content"`
}

type CreatePostHandlerResponse struct {
	ID uuid.UUID `json:"id"`
}
