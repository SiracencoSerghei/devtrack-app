package router

import (
	"net/http"
	identitytransport "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/identity/api"
	fleettransport "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/fleet/api"
	coretransport "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/core/api"
	logisticstransport "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/logistics/api"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/middleware"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/auth"
)

type Dependencies struct {
	Identity  identitytransport.Handler
	Fleet     fleettransport.Handler
	Core      coretransport.Handler
	Logistics logisticstransport.Handler
	TokenMgr  *auth.TokenManager
	CORS      string
}

func New(d Dependencies) http.Handler {
	mux := http.NewServeMux()
	authMw := middleware.Auth(d.TokenMgr)

	// System
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	// 1. IDENTITY MODULE (Автентифікація)
	mux.HandleFunc("POST /api/auth/signup", d.Identity.SignUp)
	mux.HandleFunc("POST /api/auth/login", d.Identity.Login)

	// 2. CORE MODULE (Організація/Працівники)
	mux.Handle("POST /api/core/employee", authMw(http.HandlerFunc(d.Core.OnboardEmployee)))

	// 3. FLEET MODULE (Управління автопарком/Водії)
	mux.Handle("POST /api/fleet/drivers", authMw(http.HandlerFunc(d.Fleet.CreateProfile)))
	mux.Handle("GET /api/fleet/drivers", authMw(http.HandlerFunc(d.Fleet.GetProfile)))

	// 4. LOGISTICS MODULE (Замовлення/Вантажі)
	mux.Handle("POST /api/orders", authMw(http.HandlerFunc(d.Logistics.CreateOrder)))
	mux.Handle("GET /api/orders", authMw(http.HandlerFunc(d.Logistics.GetOrder)))

	return middleware.Logging(applyCORS(mux, d.CORS))
}

func applyCORS(next http.Handler, origin string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}