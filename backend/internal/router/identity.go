package router

import (
	"net/http"
	identityapi "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/identity/api"
)

const authBase = "/api/auth"

func registerIdentityRoutes(mux *http.ServeMux, h identityapi.Handler) {
	mux.HandleFunc("POST "+authBase+"/signup", h.SignUp)
	mux.HandleFunc("POST "+authBase+"/login", h.Login)
	mux.HandleFunc("GET "+authBase+"/users", h.GetAll)
}