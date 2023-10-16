package post

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"morgan.io/internal/platform/reqctx"
	"morgan.io/internal/platform/response"
	"morgan.io/internal/post/comment"
)

type CreatePostHandlerParams struct {
	Content string `json:"content"`
}

type CreatePostHandlerResponse struct {
	ID uuid.UUID `json:"id"`
}

type CreateCommentHandlerParams struct {
	Content string `json:"content"`
}

type CreateCommentHandlerResponse struct {
	ID        uuid.UUID `json:"id"`
	PostID    uuid.UUID `json:"post_id"`
	OwnerID   uuid.UUID `json:"owner_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type service interface {
	CreatePost(ctx context.Context, params *CreatePostServiceParams) (*Post, error)
	AddLike(ctx context.Context, postID, ownerID uuid.UUID) error
	CreateComment(ctx context.Context, params *CreateCommentServiceParams) (*comment.Comment, error)
}

type handler struct {
	service service
}

func NewHandler(service service) *handler {
	return &handler{service: service}
}

func (h *handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var params CreatePostHandlerParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	userID, ok := reqctx.UserIDFromContext(r.Context())
	if !ok {
		response.Error(w, http.StatusInternalServerError, errors.New("user not found"))
		return
	}

	post, err := h.service.CreatePost(r.Context(), &CreatePostServiceParams{
		OwnerID: userID,
		Content: params.Content,
	})
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	resp := CreatePostHandlerResponse{
		ID: post.ID,
	}

	response.Success(w, http.StatusOK, resp)
}

func (h *handler) AddLike(w http.ResponseWriter, r *http.Request) {
	postID, err := uuid.Parse(mux.Vars(r)["post_id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	userID, ok := reqctx.UserIDFromContext(r.Context())
	if !ok {
		response.Error(w, http.StatusInternalServerError, errors.New("user not found"))
		return
	}

	err = h.service.AddLike(r.Context(), postID, userID)
	if err != nil {
		switch err {
		case ErrLikeAlreadyExists:
			response.Error(w, http.StatusBadRequest, err)
		case ErrPostNotFound:
			response.Error(w, http.StatusNotFound, err)
		default:
			response.Error(w, http.StatusBadRequest, err)
		}
		return
	}

	response.Success(w, http.StatusCreated, nil)
}

func (h *handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	postID, err := uuid.Parse(mux.Vars(r)["post_id"])
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	var params CreateCommentHandlerParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	userID, ok := reqctx.UserIDFromContext(r.Context())
	if !ok {
		response.Error(w, http.StatusInternalServerError, errors.New("user not found"))
		return
	}

	newComment, err := h.service.CreateComment(r.Context(), &CreateCommentServiceParams{
		PostID:  postID,
		OwnerID: userID,
		Content: params.Content,
	})
	if err != nil {
		switch err {
		case comment.ErrPostEmpty, comment.ErrEmptyContent, comment.ErrOwnerEmpty:
			response.Error(w, http.StatusBadRequest, err)
		default:
			response.Error(w, http.StatusInternalServerError, err)
		}

		return
	}

	resp := &CreateCommentHandlerResponse{
		ID:        newComment.ID,
		PostID:    newComment.PostID,
		OwnerID:   newComment.OwnerID,
		Content:   newComment.Content,
		CreatedAt: newComment.CreatedAt,
	}

	response.Success(w, http.StatusOK, resp)
}
