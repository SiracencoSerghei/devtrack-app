package bootstrap

import (
	"github.com/jackc/pgx/v5/pgxpool"

	identityapi "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/identity/api"
	identityapp "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/identity/application"
	identityinfra "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/identity/infrastructure"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/auth"

	roleapp "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/access/application"
	roleinfra "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/access/infrastructure"
)

type IdentityModule struct {
	Handler *identityapi.Handler
}

func initIdentityModule(
	pool *pgxpool.Pool,
	tm *auth.TokenManager,
) *IdentityModule {
	identityRepo := identityinfra.NewPostgresRepository(pool)

	roleRepo := roleinfra.NewPostgresRepository(pool)
	roleService := roleapp.NewService(roleRepo)

	identityService := identityapp.NewService(
		identityRepo,
		tm,
		roleService,
	)

	identityHandler := identityapi.NewHandler(identityService)

	return &IdentityModule{
		Handler: identityHandler,
	}
}