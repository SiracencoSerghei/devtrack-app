package router

import (
	"net/http"
	"encoding/json"

	"github.com/SiracencoSerghei/devtrack-app/internal/handler"
	"github.com/SiracencoSerghei/devtrack-app/internal/middleware"
	"github.com/SiracencoSerghei/devtrack-app/internal/user"
)

func New() *http.ServeMux {
	return http.NewServeMux()
}

func RegisterRootRoute(mux *http.ServeMux) {
	mux.Handle("/", middleware.Logging(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]string{"message": "server running"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})))
}

func RegisterHealthRoutes(mux *http.ServeMux, h *handler.HealthHandler) {

	mux.Handle(
		"/health",
		middleware.Logging(http.HandlerFunc(h.Health)),
	)
}

func RegisterUserRoutes(mux *http.ServeMux, h *user.Handler) {
	mux.Handle("/users", middleware.Logging(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.Create(w, r)
		case http.MethodGet:
			h.GetAll(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})))
}