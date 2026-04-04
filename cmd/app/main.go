package main

import (
    "context"
    "errors"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/SiracencoSerghei/devtrack-app/internal/health"
    "github.com/SiracencoSerghei/devtrack-app/internal/router"
    "github.com/SiracencoSerghei/devtrack-app/internal/user"
    "github.com/SiracencoSerghei/devtrack-app/pkg/httpserver"
)

func main() {
    repo := user.NewInMemoryRepository()
    svc := user.NewService(repo)

    userHandler := user.NewHandler(svc)
    healthHandler := health.NewHandler()

    r := router.New(userHandler, healthHandler)
    server := httpserver.New(":8080", r)

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

    errChan := make(chan error, 1)
    go func() {
        errChan <- server.Start()
    }()

    select {
    case err := <-errChan:
        if err != nil && !errors.Is(err, http.ErrServerClosed) {
            log.Fatalf("server error: %v", err)
        }
    case <-stop:
        log.Println("shutdown signal received")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := server.Stop(ctx); err != nil {
        log.Printf("shutdown error: %v", err)
    }

    log.Println("server stopped")
}