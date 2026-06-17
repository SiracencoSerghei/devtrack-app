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

	// Ініціалізуємо DI-контейнер
	container, err := bootstrap.NewContainer(ctx)
	if err != nil {
		slog.Error("failed to build di container", "err", err)
		os.Exit(1)
	}
	defer container.Close()

	// Запускаємо сервер в окремій горутині
	go func() {
		slog.Info("server started", "addr", container.Server.Addr)

		if err := container.Server.Start(); err != nil && err != http.ErrServerClosed {
			slog.Error("server crashed", "err", err)
			stop()
		}
	}()

	<-ctx.Done()

	// Етап Graceful Shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	slog.Info("shutdown started")

	if err := container.Server.Stop(shutdownCtx); err != nil {
		slog.Error("shutdown error", "err", err)
	}

	slog.Info("shutdown complete")
}