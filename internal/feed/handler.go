package feed

import (
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

func (h *handler) GetFeed(w http.ResponseWriter, r *http.Request) {
	userID, ok := reqctx.UserIDFromContext(r.Context())
	if !ok {
		response.Error(w, http.StatusInternalServerError, errors.New("user not found"))
		return
	}

	feed, err := h.svc.GetFeed(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	resp := &GetFeedHandlerResponse{}
	for _, p := range feed.Posts {
		resp.Posts = append(resp.Posts, &GetFeedHandlerResponsePost{
			ID:        p.ID,
			UserID:    p.OwnerID,
			Content:   p.Content,
			CreatedAt: p.CreatedAt,
		})
	}

	response.Success(w, http.StatusOK, resp)
}
