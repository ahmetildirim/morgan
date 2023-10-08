package follow

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"morgan.io/internal/platform/reqctx"
	"morgan.io/internal/platform/response"
)

type CreateFollowHandlerParams struct {
	FolloweeID uuid.UUID `json:"followee_id"`
}

type service interface {
	Follow(ctx context.Context, params *CreateFollowServiceParams) error
}

type handler struct {
	service service
}

func NewHandler(svc service) *handler {
	return &handler{
		service: svc,
	}
}

func (h *handler) CreateFollow(w http.ResponseWriter, r *http.Request) {
	var params CreateFollowHandlerParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	userID, ok := reqctx.UserIDFromContext(r.Context())
	if !ok {
		response.Error(w, http.StatusInternalServerError, errors.New("user not found"))
		return
	}

	err := h.service.Follow(r.Context(), &CreateFollowServiceParams{
		FollowerID: userID,
		FolloweeID: params.FolloweeID,
	})
	if err != nil {
		switch err {
		case ErrFolloweeNotFound, ErrAlreadyFollowing:
			response.Error(w, http.StatusBadRequest, err)
		default:
			response.Error(w, http.StatusInternalServerError, err)
		}
		return
	}

	response.Success(w, http.StatusCreated, nil)
}
