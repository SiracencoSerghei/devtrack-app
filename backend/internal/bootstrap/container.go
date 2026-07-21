package bootstrap

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/config"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/db"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/router"
	"github.com/SiracencoSerghei/devtrack-app/backend/pkg/httpserver"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/auth"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/middleware"
)

type Container struct {
	DB        *pgxpool.Pool
	Server    *httpserver.Server
	Identity  *IdentityModule
	Fleet     *FleetModule
	Core      *CoreModule
	Logistics *LogisticsModule
}

func NewContainer(ctx context.Context) (*Container, error) {
	cfg := config.LoadConfig()

	pool, err := db.Connect(ctx, cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		return nil, err
	}

	if err := db.RunMigrations(pool); err != nil {
		pool.Close()
		return nil, err
	}

	tokenManager := auth.NewTokenManager(cfg.JWTSecret)

	// Ініціалізація абсолютно ВСІХ модулів Core ERP
	identityMod := initIdentityModule(pool, tokenManager)
	fleetMod    := initFleetModule(pool)
	coreMod     := initCoreModule(pool)
	logisticsMod := initLogisticsModule(pool)

r := router.New(router.Dependencies{
		Config: router.Config{
			CORS: middleware.CORSConfig{
				AllowedOrigin: cfg.CORSOrigin,
			},
		},
		Identity:  *identityMod.Handler,
		Fleet:     *fleetMod.Handler,
		Core:      *coreMod.Handler,
		Logistics: *logisticsMod.Handler,
		TokenMgr:  tokenManager,
	})

	server := httpserver.New(":"+cfg.Port, r)

	return &Container{
		DB:        pool,
		Server:    server,
		Identity:  identityMod,
		Fleet:     fleetMod,
		Core:      coreMod,
		Logistics: logisticsMod,
	}, nil
}

func (c *Container) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}