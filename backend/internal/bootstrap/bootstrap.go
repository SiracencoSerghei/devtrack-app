package bootstrap

import (
	"context"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/auth"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/config"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/db"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/health"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/role"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/router"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/user"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/driver"
	"github.com/SiracencoSerghei/devtrack-app/backend/pkg/httpserver"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Server *httpserver.Server
	pool   *pgxpool.Pool
}

func Initialize(ctx context.Context) (*App, error) {
	cfg := config.LoadConfig()

	pool, err := db.Connect(ctx, cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		return nil, err
	}

	if err := db.RunMigrations(ctx, pool); err != nil {
		return nil, err
	}

	tokenManager := auth.NewTokenManager(cfg.JWTSecret)

	roleRepo := role.NewPostgresRepository(pool)
	userRepo := user.NewPostgresRepository(pool)
	driverRepo := driver.NewPostgresRepository(pool)

	userService := user.NewService(userRepo, roleRepo, tokenManager)
	driverService := driver.NewService(driverRepo)

	userHandler := user.NewHandler(userService)
	driverHandler := driver.NewHandler(driverService)
	healthHandler := health.NewHandler()

	r := router.New(userHandler, healthHandler, driverHandler, tokenManager, cfg.CORSOrigin)

	return &App{
		Server: httpserver.New(":"+cfg.Port, r),
		pool:   pool,
	}, nil
}

func (a *App) Close() {
	if a.pool != nil {
		a.pool.Close()
	}
}