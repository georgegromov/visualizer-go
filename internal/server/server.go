package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"visualizer-go/internal/config"
)

type Server struct {
	log        *slog.Logger
	config     config.Server
	httpServer *http.Server
}

func New(log *slog.Logger, config config.Server, handler http.Handler) *Server {

	httpServer := &http.Server{
		Addr:           ":" + config.Port,
		Handler:        handler,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		IdleTimeout:    config.IdleTimeout,
		MaxHeaderBytes: config.MaxHeaderMegabytes << 20,
	}

	return &Server{
		log:        log,
		config:     config,
		httpServer: httpServer,
	}
}

func (s *Server) MustRun() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}

func (s *Server) Run() error {
	const op = "server.Run"

	log := s.log.With(slog.String("op", op), slog.String("port", s.config.Port))

	log.Info("starting http server...", slog.String("addr", s.config.Host+":"+s.config.Port))

	if err := s.httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	const op = "server.Shutdown"

	log := s.log.With(slog.String("op", op))

	log.Info("stopping http server...")

	return s.httpServer.Shutdown(ctx)
}
