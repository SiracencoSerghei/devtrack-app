package bootstrap

import (
	"github.com/jackc/pgx/v5/pgxpool"
	accessapp "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/access/application"
	accessinfra "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/access/infrastructure"
	identityapp "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/identity/application"
	identityinfra "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/identity/infrastructure"
	identitytransport "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/identity/api"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/auth"
)

type IdentityModule struct {
	Handler *identitytransport.Handler
}

func initIdentityModule(pool *pgxpool.Pool, tm *auth.TokenManager) *IdentityModule {
	accessRepo := accessinfra.NewPostgresRepository(pool)
	accessService := accessapp.NewService(accessRepo)

	identityRepo := identityinfra.NewPostgresRepository(pool)
	identityService := identityapp.NewService(identityRepo, tm, accessService)
	identityHandler := identitytransport.NewHandler(identityService)

	return &IdentityModule{Handler: identityHandler}
}