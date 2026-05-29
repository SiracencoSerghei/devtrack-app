package router

import (
	"net/http"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/user"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/health"
	
)

func New(u *user.Handler, h *health.Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/signup", u.SignUp)
	mux.HandleFunc("POST /api/login", u.Login)
	
	mux.HandleFunc("GET /api/users", u.GetAll)
	
	mux.HandleFunc("GET /health", h.HealthCheck)

	return applyCORS(mux)
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