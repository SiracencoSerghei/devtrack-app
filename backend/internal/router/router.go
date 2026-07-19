package router

import (
	"net/http"
	identityapi "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/identity/api"
	fleetapi "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/fleet/api"
	coreapi "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/core/api"
	logisticsapi "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/logistics/api"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/middleware"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/auth"
)

type Config struct {
	CORS middleware.CORSConfig
}

type Dependencies struct {
	Config    Config
	Identity  identityapi.Handler
	Fleet     fleetapi.Handler
	Core      coreapi.Handler
	Logistics logisticsapi.Handler
	TokenMgr  *auth.TokenManager
}

func New(d Dependencies) http.Handler {
	mux := http.NewServeMux()
	authMw := middleware.Auth(d.TokenMgr)

	// Декларативна реєстрація ізольованих модулів маршрутизації
	registerSystemRoutes(mux)
	registerIdentityRoutes(mux, d.Identity)
	registerCoreRoutes(mux, authMw, d.Core)
	registerFleetRoutes(mux, authMw, d.Fleet)
	registerLogisticsRoutes(mux, authMw, d.Logistics)

	// Чистий ланцюжок middleware
	corsMw := middleware.CORS(d.Config.CORS)
	
	return middleware.Logging(corsMw(mux))
}