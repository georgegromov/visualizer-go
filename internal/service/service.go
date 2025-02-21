package service

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/models"
	"visualizer-go/internal/repository"
)

type (
	Template interface {
		GetAll(ctx context.Context) ([]models.Template, error)
		GetByID(ctx context.Context, templateID uuid.UUID) (models.Template, error)
		Create(ctx context.Context, dto dto.TemplateCreateDto) error
		Update(ctx context.Context, templateID uuid.UUID, dto dto.TemplateUpdateDto) error
	}

	Deps struct {
		Repo *repository.Repository
	}

	Service struct {
		Template
	}
)

func New(log *slog.Logger, deps Deps) *Service {
	return &Service{
		Template: NewTemplateService(log, deps.Repo.Template),
	}
}
