package bootstrap

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/db"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/health"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/role"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/router"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/user"
	"github.com/SiracencoSerghei/devtrack-app/backend/internal/driver"
	"github.com/jackc/pgx/v5/pgxpool"
)

// App racchiude tutte le risorse core dell'ERP per la gestione del ciclo di vita
type App struct {
	Server *http.Server
	pool   *pgxpool.Pool
}

// Initialize centralizza la creazione di repository, service, handler e router (Dependency Injection)
func Initialize(ctx context.Context) (*App, error) {
	// 1. Connessione al Database Pool
	pool, err := db.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	// 2. Inizializzazione Strato Repository
	roleRepo := role.NewPostgresRepository(pool)
	userRepo := user.NewPostgresRepository(pool)
	driverRepo := driver.NewPostgresRepository(pool)

	// 3. Inizializzazione Strato Business Logic (Service)
	userService := user.NewService(userRepo, roleRepo)
	driverService := driver.NewService(driverRepo)

	// 4. Inizializzazione Strato HTTP (Handler)
	userHandler := user.NewHandler(userService)
	healthHandler := health.NewHandler()
	driverHandler := driver.NewHandler(driverService)

	// 5. Configurazione del Router con tutti gli endpoint dell'ERP
	appRouter := router.New(userHandler, healthHandler, driverHandler)

	// 6. Configurazione dell'HTTP Server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      appRouter,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	return &App{
		Server: server,
		pool:   pool,
	}, nil
}

// Close si occupa di liberare le risorse in modo sicuro all'arresto dell'applicazione
func (a *App) Close() {
	if a.pool != nil {
		a.pool.Close()
	}
}