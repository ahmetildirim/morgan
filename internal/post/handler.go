package post

import (
	"encoding/json"
	"errors"
	"net/http"

	"morgan.io/internal/platform/reqctx"
	"morgan.io/internal/platform/response"
)

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
