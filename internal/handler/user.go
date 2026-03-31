package handler

import (
	"encoding/json"
	"net/http"

	"github.com/SiracencoSerghei/devtrack-app/internal/service"
	"github.com/SiracencoSerghei/devtrack-app/internal/httpx"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

type UserResponse struct {
	Data string `json:"data"`
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		httpx.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user, err := h.service.GetUser(r.Context())
	if err != nil {
		httpx.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := UserResponse{
		Data: user,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		httpx.WriteError(w, http.StatusInternalServerError, "encoding error")
	}
}