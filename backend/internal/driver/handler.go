package driver

import (
	"encoding/json"
	"net/http"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/middleware"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

type createDriverReq struct {
	LicenseNumber string `json:"license_number"`
	Phone         string `json:"phone"`
}

type updateStatusReq struct {
	Status string `json:"status"`
}

func (h *Handler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	var req createDriverReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	claims, ok := middleware.GetUser(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	d, err := h.svc.CreateProfile(r.Context(), claims.UserID, req.LicenseNumber, req.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(d)
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id query parameter", http.StatusBadRequest)
		return
	}

	d, err := h.svc.GetProfileByUserID(r.Context(), userID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(d)
}