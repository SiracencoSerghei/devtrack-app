package api

import (
	"encoding/json"
	"net/http"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/core/application"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/core/domain"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/httpx"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/middleware"
)

type Handler struct {
	svc *application.Service
}

func NewHandler(svc *application.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) OnboardEmployee(w http.ResponseWriter, r *http.Request) {
	var emp domain.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		httpx.WriteJSON(w, http.StatusBadRequest, httpx.ErrorResponse{Error: "Invalid JSON"})
		return
	}

	claims, ok := middleware.GetUser(r.Context())
	if !ok {
		httpx.WriteJSON(w, http.StatusUnauthorized, httpx.ErrorResponse{Error: "Unauthorized"})
		return
	}
	emp.UserID = claims.UserID

	res, err := h.svc.RegisterEmployee(r.Context(), emp)
	if err != nil {
		httpx.WriteJSON(w, http.StatusBadRequest, httpx.ErrorResponse{Error: err.Error()})
		return
	}
	httpx.WriteJSON(w, http.StatusCreated, res)
}