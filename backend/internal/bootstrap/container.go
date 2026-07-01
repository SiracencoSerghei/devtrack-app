package bootstrap

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/config"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/db"
	identityinfra "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/identity/infrastructure"
	identityapp "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/identity/application"
	identitytransport "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/identity/transport"
	logisticsinfra "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/logistics/infrastructure"
	logisticsapp "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/logistics/application"
	logisticstransport "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/logistics/transport"
	accessinfra "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/access/infrastructure"
	accessapp "github.com/SiracencoSerghei/devtrack-app/backend/internal/contexts/access/application"
	
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/router"
	"github.com/SiracencoSerghei/devtrack-app/backend/pkg/httpserver"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/shared/auth"
)

type Container struct {
	DB        *pgxpool.Pool
	Server    *httpserver.Server
	Identity  identitytransport.Handler
	Logistics logisticstransport.Handler
}

func NewContainer(ctx context.Context) (*Container, error) {
	cfg := config.LoadConfig()

	// Передаємо параметри конфігурації у твій Connect метод
	pool, err := db.Connect(ctx, cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		return nil, err
	}

	// Запуск міграцій на старті додатка
	if err := db.RunMigrations(pool); err != nil {
		return nil, err
	}

	tokenManager := auth.NewTokenManager(cfg.JWTSecret)
	accessRepo := accessinfra.NewPostgresRepository(pool)
	accessService := accessapp.NewService(accessRepo)

	identityRepo := identityinfra.NewPostgresRepository(pool)


	identityService := identityapp.NewService(
		identityRepo,
		tokenManager,
		accessService,
	)

	identityHandler := identitytransport.NewHandler(identityService)

	logisticsRepo := logisticsinfra.NewPostgresRepository(pool)
	logisticsService := logisticsapp.NewService(logisticsRepo)
	logisticsHandler := logisticstransport.NewHandler(logisticsService)


	r := router.New(
		router.Dependencies{
			Identity:   *identityHandler,
			Logistics:  *logisticsHandler,
			TokenMgr:   tokenManager,
			CORS:       cfg.CORSOrigin,
		},
	)

	server := httpserver.New(":"+cfg.Port, r)

	return &Container{
		DB:        pool,
		Server:    server,
		Identity:  *identityHandler,
		Logistics: *logisticsHandler,
	}, nil
}

func (c *Container) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}