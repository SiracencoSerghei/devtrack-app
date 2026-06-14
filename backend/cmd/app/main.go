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
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app, err := bootstrap.Initialize(ctx)
	if err != nil {
		slog.Error("bootstrap failed", "err", err)
		os.Exit(1)
	}
	defer app.Close()

	go func() {
		slog.Info("server started", "addr", app.Server.Addr)

		err := app.Server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			slog.Error("server crashed", "err", err)
			stop()
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	slog.Info("shutdown started")

	if err := app.Server.Shutdown(shutdownCtx); err != nil {
		slog.Error("shutdown error", "err", err)
	}

	slog.Info("shutdown complete")
}