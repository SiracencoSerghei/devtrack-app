package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SiracencoSerghei/devtrack-app/backend/internal/bootstrap"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	container, err := bootstrap.NewContainer(ctx)
	if err != nil {
		slog.Error("container build failed", "err", err)
		os.Exit(1)
	}
	defer container.Close()

	server := container.Server

	go func() {
		slog.Info("server started", "addr", server.Addr)
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			stop()
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = server.Shutdown(shutdownCtx)
}