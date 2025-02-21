package main

import (
	"context"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
	"visualizer-go/internal/handler"
	"visualizer-go/internal/lib/config"
	"visualizer-go/internal/lib/db/postgres"
	"visualizer-go/internal/lib/server"
	"visualizer-go/internal/repository"
	"visualizer-go/internal/service"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	log.Info("initializing server...", slog.String("address", cfg.Server.Host+":"+cfg.Server.Port))
	log.Debug("logger debug mode enabled")

	db := postgres.MustConnect(log, cfg.Database)
	repo := repository.New(log, db)
	svc := service.New(log, service.Deps{
		Repo: repo,
	})
	h := handler.New(log, svc)

	srv := server.New(log, cfg.Server, h.Init())

	go func() {
		srv.MustRun()
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	sign := <-stop

	log.Info("gracefully stopping application...", slog.String("signal", sign.String()))

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Error("error occurred while stopping http server: %s", err.Error())
	}

	log.Info("server successfully stopped")

	if err := db.Close(); err != nil {
		log.Error("error occurred while closing database: %s", err.Error())
	}

	log.Info("postgres successfully closed")

	log.Info("application gracefully stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
