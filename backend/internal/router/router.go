package router

import (
	"net/http"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/health"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/middleware"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/user"
)

func New(u *user.Handler, h *health.Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"app":"DevTrack API","status":"running"}`))
	})

	mux.HandleFunc("POST /api/signup", u.SignUp)
	mux.HandleFunc("POST /api/login", u.Login)
	mux.HandleFunc("GET /health", h.HealthCheck)

	mux.Handle("GET /api/users",
		middleware.Auth(http.HandlerFunc(u.GetAll)),
	)

	return middleware.Logging(applyCORS(mux))
}

func applyCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}