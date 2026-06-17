package bootstrap

import (
	"context"
	"fmt"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/auth"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/config"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/db"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/driver"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/health"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/role"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/router"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/user"
	"github.com/SiracencoSerghei/devtrack-app/backend/pkg/httpserver"
	
	"github.com/jackc/pgx/v5/pgxpool"
)

// Container інкапсулює всі довгоживучі інфраструктурні та бізнес-компоненти
type Container struct {
	Config *config.Config
	Pool   *pgxpool.Pool
	Server *httpserver.Server
}

// NewContainer будує та валідує весь граф залежностей (DI)
func NewContainer(ctx context.Context) (*Container, error) {
	cfg := config.LoadConfig()

	// 1. Інфраструктурний шар
	pool, err := db.Connect(ctx, cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		return nil, fmt.Errorf("db connection failed: %w", err)
	}

	if err := db.RunMigrations(pool); err != nil {
		pool.Close()
		return nil, fmt.Errorf("migrations failed: %w", err)
	}

	tokenManager := auth.NewTokenManager(cfg.JWTSecret)

	// 2. Репозиторії
	roleRepo := role.NewPostgresRepository(pool)
	userRepo := user.NewPostgresRepository(pool)
	driverRepo := driver.NewPostgresRepository(pool)

	// 3. Сервіси
	userService := user.NewService(userRepo, roleRepo, tokenManager)
	driverService := driver.NewService(driverRepo)

	// 4. Хендлери
	userHandler := user.NewHandler(userService)
	driverHandler := driver.NewHandler(driverService)
	healthHandler := health.NewHandler()

	// 5. Роутер та Сервер
	r := router.New(userHandler, healthHandler, driverHandler, tokenManager, cfg.CORSOrigin)
	server := httpserver.New(":"+cfg.Port, r)

	return &Container{
		Config: cfg,
		Pool:   pool,
		Server: server,
	}, nil
}

// Close безпечно звільняє ресурси контейнера при завершенні роботи
func (c *Container) Close() {
	if c.Pool != nil {
		c.Pool.Close()
	}
}