package post

import "github.com/google/uuid"

type CreatePostServiceParams struct {
	OwnerID uuid.UUID
	Content string
}
