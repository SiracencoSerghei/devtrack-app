package router

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"

    "github.com/SiracencoSerghei/devtrack-app/internal/health"
    "github.com/SiracencoSerghei/devtrack-app/internal/user"
)

func New() *chi.Mux {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    return r
}

func RegisterRootRoute(r *chi.Mux) {
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        _ = json.NewEncoder(w).Encode(map[string]string{"message": "server running"})
    })
}

func RegisterHealthRoutes(r *chi.Mux, h *health.Handler) {
    r.Get("/health", h.Check)
}

func RegisterUserRoutes(r *chi.Mux, h *user.Handler) {
    r.Post("/users", h.Create)
    r.Get("/users", h.GetAll)
}