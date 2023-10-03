package like

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
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

	err = h.svc.Create(r.Context(), postID, userID)
	if err != nil {
		switch err {
		case ErrPostNotFound:
			response.Error(w, http.StatusBadRequest, err)
		default:
			response.Error(w, http.StatusInternalServerError, err)
		}
		return
	}

	response.Success(w, http.StatusCreated, nil)
}
