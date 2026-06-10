package driver

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

type createDriverReq struct {
	UserID        string `json:"user_id"`
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

	d, err := h.svc.CreateProfile(r.Context(), req.UserID, req.LicenseNumber, req.Phone)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
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