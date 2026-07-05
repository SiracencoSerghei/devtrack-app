package api

import (
	"encoding/json"
	"net/http"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/logistics/application"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/httpx"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/middleware"
)

type createOrderReq struct {
	PickupAddress   string `json:"pickup_address"`
	DeliveryAddress string `json:"delivery_address"`
}

type Handler struct {
	svc *application.Service
}

func NewHandler(svc *application.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req createOrderReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteJSON(w, http.StatusBadRequest, httpx.ErrorResponse{Error: "Invalid JSON format"})
		return
	}

	claims, ok := middleware.GetUser(r.Context())
	if !ok {
		httpx.WriteJSON(w, http.StatusUnauthorized, httpx.ErrorResponse{Error: "Unauthorized"})
		return
	}

	// Використовуємо UserID користувача як CustomerID для спрощення MVP
	o, err := h.svc.CreateOrder(r.Context(), claims.UserID, req.PickupAddress, req.DeliveryAddress)
	if err != nil {
		httpx.WriteError(w, err)
		return
	}

	httpx.WriteJSON(w, http.StatusCreated, o)
}

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		httpx.WriteJSON(w, http.StatusBadRequest, httpx.ErrorResponse{Error: "Missing id parameter"})
		return
	}

	o, err := h.svc.GetOrder(r.Context(), id)
	if err != nil {
		httpx.WriteError(w, err)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, o)
}