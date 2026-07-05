package bootstrap

import (
	"github.com/jackc/pgx/v5/pgxpool"
	logisticsapp "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/logistics/application"
	logisticsinfra "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/logistics/infrastructure"
	logisticsapi "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/logistics/api"
)

type LogisticsModule struct {
	Handler *logisticsapi.Handler
}

func initLogisticsModule(pool *pgxpool.Pool) *LogisticsModule {
	repo := logisticsinfra.NewPostgresRepository(pool)
	svc := logisticsapp.NewService(repo)
	handler := logisticsapi.NewHandler(svc)
	return &LogisticsModule{Handler: handler}
}