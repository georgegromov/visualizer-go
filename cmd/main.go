package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "visualizer-go/docs"
	"visualizer-go/internal/api/server"
	"visualizer-go/internal/config"
	"visualizer-go/internal/database/postgres"
	"visualizer-go/pkg/logger"

	_ "github.com/lib/pq"
)

// @title Visualizer REST API
// @version 1.0
// @description A REST API in Go using Gin framework

//	@contact.name
//	@contact.url
//	@contact.email

//	@license.name
//	@license.url

//	@host		localhost:8888
//	@BasePath	/api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format **Bearer &lt;token&gt;**

func main() {
	cfg := config.MustLoad()

	log := logger.NewLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	log.Info("starting application...", slog.String("address", cfg.Server.Host+":"+cfg.Server.Port))
	log.Debug("logger debug mode enabled")

	pgdb := postgres.MustConnect(log, cfg.Database)

	httpServer := server.New(log, cfg, pgdb)
	httpServer.Register()

	go httpServer.MustStart()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	sign := <-stop

	log.Info("gracefully stopping application...", slog.String("signal", sign.String()))

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := httpServer.Stop(ctx); err != nil {
		log.Error("error occurred while stopping http server", slog.String("error", err.Error()))
	}

	log.Info("server successfully stopped")

	if err := pgdb.Close(); err != nil {
		log.Error("error occurred while closing database", slog.String("error", err.Error()))
	}

	log.Info("postgres successfully closed")

	log.Info("application gracefully stopped")
}
