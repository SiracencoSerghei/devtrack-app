package bootstrap

import (
	"github.com/jackc/pgx/v5/pgxpool"
	fleetapp "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/fleet/application"
	fleetinfra "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/fleet/infrastructure"
	fleettransport "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/fleet/api"
)

type FleetModule struct {
	Handler *fleettransport.Handler
}

func initFleetModule(pool *pgxpool.Pool) *FleetModule {
	fleetRepo := fleetinfra.NewPostgresRepository(pool)
	fleetService := fleetapp.NewService(fleetRepo)
	fleetHandler := fleettransport.NewHandler(fleetService)

	return &FleetModule{Handler: fleetHandler}
}