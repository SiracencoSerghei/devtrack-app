package health

import (
    "net/http"

    "github.com/SiracencoSerghei/devtrack-app/internal/httpx"
)

type Handler struct{}

func NewHandler() *Handler {
    return &Handler{}
}

func (h *Handler) Check(w http.ResponseWriter, r *http.Request) {
    httpx.WriteJSON(w, http.StatusOK, map[string]string{
        "status": "ok",
    })
}