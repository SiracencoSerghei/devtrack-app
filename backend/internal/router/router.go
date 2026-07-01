package router

import (
	"net/http"
	identitytransport "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/identity/transport"
	logisticstransport "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/logistics/transport"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/middleware"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/auth"
)

type Dependencies struct {
	Identity   identitytransport.Handler
	Logistics  logisticstransport.Handler
	TokenMgr   *auth.TokenManager
	CORS       string
}

func New(d Dependencies) http.Handler {
	mux := http.NewServeMux()

	// public
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	// identity
	mux.HandleFunc("POST /api/signup", d.Identity.SignUp)
	mux.HandleFunc("POST /api/login", d.Identity.Login)

	// protected
	authMw := middleware.Auth(d.TokenMgr)
	mux.Handle("POST /api/drivers",
		authMw(http.HandlerFunc(d.Logistics.CreateProfile)),
	)
	mux.Handle("GET /api/drivers",
		authMw(http.HandlerFunc(d.Logistics.GetProfile)),
	)

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