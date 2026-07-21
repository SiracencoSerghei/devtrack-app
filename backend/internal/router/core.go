package router

import (
	"net/http"
	coreapi "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/core/api"
)

const coreBase = "/api/core"

func registerCoreRoutes(mux *http.ServeMux, authMw func(http.Handler) http.Handler, h coreapi.Handler) {
	mux.Handle("POST "+coreBase+"/employee", authMw(http.HandlerFunc(h.OnboardEmployee)))
}