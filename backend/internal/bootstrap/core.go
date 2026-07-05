package bootstrap

import (
	"github.com/jackc/pgx/v5/pgxpool"
	coreapp "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/core/application"
	coreinfra "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/core/infrastructure"
	coreapi "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/core/api"
)

type CoreModule struct {
	Handler *coreapi.Handler
}

func initCoreModule(pool *pgxpool.Pool) *CoreModule {
	repo := coreinfra.NewPostgresRepository(pool)
	svc := coreapp.NewService(repo)
	handler := coreapi.NewHandler(svc)
	return &CoreModule{Handler: handler}
}