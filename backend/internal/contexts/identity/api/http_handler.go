package transport

import (
	"encoding/json"
	"net/http"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/httpx"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/identity/application"
)

type signUpReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Handler struct {
	svc application.ServiceInterface // Зав'язано на інтерфейс
}

func NewHandler(svc application.ServiceInterface) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req signUpReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteJSON(w, http.StatusBadRequest, httpx.ErrorResponse{Error: "Invalid JSON format"})
		return
	}

	u, err := h.svc.SignUp(r.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		httpx.WriteError(w, err) // Авто-мапінг через новий слой помилок!
		return
	}

	httpx.WriteJSON(w, http.StatusCreated, u)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteJSON(w, http.StatusBadRequest, httpx.ErrorResponse{Error: "Invalid JSON format"})
		return
	}

	token, u, err := h.svc.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		httpx.WriteError(w, err)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, map[string]any{"token": token, "user": u})
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.svc.GetAll(r.Context())
	if err != nil {
		httpx.WriteError(w, err)
		return
	}
	httpx.WriteJSON(w, http.StatusOK, users)
}