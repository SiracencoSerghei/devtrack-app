package router

import (
	"net/http"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/health"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/middleware"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/user"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/driver"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/auth"
)

func New(u *user.Handler, h *health.Handler, d *driver.Handler, tm *auth.TokenManager, corsOrigin string) http.Handler {
	mux := http.NewServeMux()

	authMiddleware := middleware.Auth(tm)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"app":"DevTrack API","status":"running"}`))
	})

	mux.HandleFunc("POST /api/signup", u.SignUp)
	mux.HandleFunc("POST /api/login", u.Login)
	mux.HandleFunc("GET /health", h.HealthCheck)

	mux.Handle("POST /api/drivers", authMiddleware(http.HandlerFunc(d.CreateProfile)))
	mux.Handle("GET /api/drivers", authMiddleware(http.HandlerFunc(d.GetProfile)))

	mux.Handle("GET /api/users", authMiddleware(http.HandlerFunc(u.GetAll)))

	return middleware.Logging(applyCORS(mux, corsOrigin))
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