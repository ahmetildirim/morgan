package follow

import (
	"encoding/json"
	"errors"
	"net/http"

	"morgan.io/internal/platform/reqctx"
	"morgan.io/internal/platform/response"
)

type handler struct {
	svc service
}

func NewHandler(svc service) *handler {
	return &handler{
		svc: svc,
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

	err := h.svc.Follow(r.Context(), &CreateFollowServiceParams{
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
