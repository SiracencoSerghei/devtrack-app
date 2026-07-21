package router

import (
	"net/http"
	logisticsapi "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/logistics/api"
)

const logisticsBase = "/api/orders"

func registerLogisticsRoutes(mux *http.ServeMux, authMw func(http.Handler) http.Handler, h logisticsapi.Handler) {
	mux.Handle("POST "+logisticsBase, authMw(http.HandlerFunc(h.CreateOrder)))
	mux.Handle("GET "+logisticsBase, authMw(http.HandlerFunc(h.GetOrder)))
}