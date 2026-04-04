package router

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi/v5"
    chimiddleware "github.com/go-chi/chi/v5/middleware"

    appmiddleware "github.com/SiracencoSerghei/devtrack-app/internal/middleware"
    "github.com/SiracencoSerghei/devtrack-app/internal/health"
    "github.com/SiracencoSerghei/devtrack-app/internal/user"
)

func New(userHandler *user.Handler, healthHandler *health.Handler) *chi.Mux {
    r := chi.NewRouter()

    r.Use(chimiddleware.Recoverer)
    r.Use(chimiddleware.RequestID)
    r.Use(appmiddleware.Logging)

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        _ = json.NewEncoder(w).Encode(map[string]string{
            "message": "server running",
        })
    })

    r.Get("/health", healthHandler.Check)

    r.Route("/users", func(r chi.Router) {
        r.Post("/", userHandler.Create)
        r.Get("/", userHandler.GetAll)
    })

    return r
}