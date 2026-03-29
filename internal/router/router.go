package router

import (
	"net/http"

	"github.com/SiracencoSerghei/devtrack-app/internal/handler"
	"github.com/SiracencoSerghei/devtrack-app/internal/middleware"
)

func New() *http.ServeMux {
	return http.NewServeMux()
}

func RegisterHealthRoutes(mux *http.ServeMux, h *handler.HealthHandler) {

	mux.Handle(
		"/health",
		middleware.Logging(http.HandlerFunc(h.Health)),
	)
}

func RegisterUserRoutes(mux *http.ServeMux, h *handler.UserHandler) {

	mux.Handle(
		"/users",
		middleware.Logging(http.HandlerFunc(h.GetUsers)),
	)
}