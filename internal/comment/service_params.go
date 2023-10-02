package comment

import "github.com/google/uuid"

type CreateCommentServiceParams struct {
	PostID  uuid.UUID
	OwnerID uuid.UUID
	Content string
}
