package auth

import (
	"encoding/json"
	"net/http"

	"morgan.io/internal/platform/response"
)

type Handler struct {
	service service
}

func NewHandler(service service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var params LoginHandlerParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	token, err := h.service.Login(r.Context(), &LoginServiceParams{
		Email:    params.Email,
		Password: params.Password,
	})
	if err != nil {
		switch err {
		default:
			response.Error(w, http.StatusUnauthorized, err)
		}
		return
	}

	resp := LoginHandlerResponse{
		Token: token,
	}

	response.Success(w, http.StatusOK, resp)
}
