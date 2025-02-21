package service

import (
	"log/slog"
	"visualizer-go/internal/repository"
)

type (
	Deps struct {
		Repo *repository.Repository
	}

	Service struct {
	}
)

func New(log *slog.Logger, deps Deps) *Service {
	return &Service{}
}
