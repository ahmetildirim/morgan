package user

import (
	"context"
	"encoding/json"
	"net/http"

	"morgan.io/internal/platform/response"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	ID string `json:"id"`
}

type service interface {
	CreateUser(ctx context.Context, params *CreateServiceParams) (*User, error)
	Authenticate(ctx context.Context, email string, password string) (*User, error)
}

type handler struct {
	svc service
}

func NewHandler(svc service) *handler {
	return &handler{
		svc: svc,
	}
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var params CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.svc.CreateUser(r.Context(), &CreateServiceParams{
		Email:    params.Email,
		Password: params.Password,
	})

	if err != nil {
		switch err {
		case
			ErrEmailAlreadyExists,
			ErrInvalidEmail,
			ErrInvalidPassword:
			response.Error(w, http.StatusBadRequest, err)
		default:
			response.Error(w, http.StatusInternalServerError, err)
		}
		return
	}

	resp := CreateUserResponse{
		ID: user.ID.String(),
	}

	response.Success(w, http.StatusCreated, resp)
}
