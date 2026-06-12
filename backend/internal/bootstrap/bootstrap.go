package bootstrap

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/config"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/auth"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/db"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/health"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/role"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/router"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/user"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/driver"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Server *http.Server
	pool   *pgxpool.Pool
}

func Initialize(ctx context.Context) (*App, error) {
	// 0. Завантажуємо конфігурацію
	cfg := config.LoadConfig()

	// 1. Connessione al Database Pool (передаємо параметри з cfg)
	pool, err := db.Connect(ctx, cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	// Ініціалізуємо менеджер токенів
	tokenManager := auth.NewTokenManager(cfg.JWTSecret)

	// 2. Inizializzazione Strato Repository
	roleRepo := role.NewPostgresRepository(pool)
	userRepo := user.NewPostgresRepository(pool)
	driverRepo := driver.NewPostgresRepository(pool)

	// 3. Inizializzazione Strato Business Logic (передаємо tokenManager туди, де він потрібен)
	userService := user.NewService(userRepo, roleRepo, tokenManager)
	driverService := driver.NewService(driverRepo)

	// 4. Inizializzazione Strato HTTP
	userHandler := user.NewHandler(userService)
	healthHandler := health.NewHandler()
	driverHandler := driver.NewHandler(driverService)

	// 5. Configurazione del Router (сюди теж передамо tokenManager для Middleware)
	appRouter := router.New(userHandler, healthHandler, driverHandler, tokenManager, cfg.CORSOrigin)

	// 6. Configurazione dell'HTTP Server (порт беремо з конфігу)
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      appRouter,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	return &App{
		Server: server,
		pool:   pool,
	}, nil
}

func (a *App) Close() {
	if a.pool != nil {
		a.pool.Close()
	}
}