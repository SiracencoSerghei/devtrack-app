package transport

import (
	"encoding/json"
	"net/http"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/httpx"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/middleware"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/logistics/application"
)

type createDriverReq struct {
	LicenseNumber string `json:"license_number"`
	Phone         string `json:"phone"`
}

type Handler struct {
	svc application.ServiceInterface
}

func NewHandler(svc application.ServiceInterface) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	var req createDriverReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteJSON(w, http.StatusBadRequest, httpx.ErrorResponse{Error: "Invalid JSON format"})
		return
	}

	claims, ok := middleware.GetUser(r.Context())
	if !ok {
		httpx.WriteJSON(w, http.StatusUnauthorized, httpx.ErrorResponse{Error: "Unauthorized"})
		return
	}

	d, err := h.svc.CreateProfile(r.Context(), claims.UserID, req.LicenseNumber, req.Phone)
	if err != nil {
		httpx.WriteError(w, err)
		return
	}

	httpx.WriteJSON(w, http.StatusCreated, d)
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		httpx.WriteJSON(w, http.StatusBadRequest, httpx.ErrorResponse{Error: "Missing user_id parameter"})
		return
	}

	d, err := h.svc.GetProfileByUserID(r.Context(), userID)
	if err != nil {
		httpx.WriteError(w, err)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, d)
}