package comment

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"morgan.io/internal/platform/reqctx"
	"morgan.io/internal/platform/response"
)

type handler struct {
	service service
}

func NewHandler(svc service) *handler {
	return &handler{
		service: svc,
	}
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
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

	comment, err := h.service.CreateComment(r.Context(), &CreateCommentServiceParams{
		PostID:  postID,
		OwnerID: userID,
		Content: params.Content,
	})
	if err != nil {
		switch err {
		case ErrPostEmpty, ErrEmptyContent, ErrOwnerEmpty:
			response.Error(w, http.StatusBadRequest, err)
		default:
			response.Error(w, http.StatusInternalServerError, err)
		}

		return
	}

	resp := &CreateCommentHandlerResponse{
		ID:        comment.ID,
		PostID:    comment.PostID,
		OwnerID:   comment.OwnerID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
	}

	response.Success(w, http.StatusOK, resp)
}
