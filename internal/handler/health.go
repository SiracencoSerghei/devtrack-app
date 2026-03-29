package handler

import (
	"encoding/json"
	"net/http"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthResponse struct {
	Status string `json:"status"`
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	resp := HealthResponse{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		WriteError(w, http.StatusInternalServerError, "encoding error")
	}
}