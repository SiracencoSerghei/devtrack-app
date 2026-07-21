package router

import (
	"net/http"
	fleetapi "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/fleet/api"
)

const fleetBase = "/api/fleet"

func registerFleetRoutes(mux *http.ServeMux, authMw func(http.Handler) http.Handler, h fleetapi.Handler) {
	mux.Handle("POST "+fleetBase+"/drivers", authMw(http.HandlerFunc(h.CreateProfile)))
	mux.Handle("GET "+fleetBase+"/drivers", authMw(http.HandlerFunc(h.GetProfile)))
}