package service

import (
	"context"
	"log/slog"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/models"
	"visualizer-go/internal/repository"

	"github.com/google/uuid"
)

type (
	User interface {
		Login(ctx context.Context, dto dto.UserLoginDto) (models.User, string, error)
		GetByID(ctx context.Context, userID uuid.UUID) (models.User, error)
		GetByUsername(ctx context.Context, username string) (models.User, error)
		Create(ctx context.Context, dto dto.UserCreateDto) error
		Update(ctx context.Context, userID uuid.UUID, dto dto.UserUpdateDto) error
	}
	Template interface {
		GetAll(ctx context.Context) ([]models.Template, error)
		GetByID(ctx context.Context, templateID uuid.UUID) (models.Template, error)
		Create(ctx context.Context, dto dto.TemplateCreateDto) (uuid.UUID, error)
		Update(ctx context.Context, templateID uuid.UUID, dto dto.TemplateUpdateDto) error
	}

	Canvas interface {
		Create(ctx context.Context, dto dto.CanvasCreateDto) error
		GetCanvasesByTemplateID(ctx context.Context, templateID uuid.UUID) ([]models.Canvas, error)
		Update(ctx context.Context, canvasID uuid.UUID, dto dto.CanvasUpdateDto) error
		Delete(ctx context.Context, canvasID uuid.UUID) error
	}

	Chart interface {
		GetByCanvasID(ctx context.Context, canvasID uuid.UUID) ([]models.Chart, error)
		Create(ctx context.Context, dto dto.ChartCreateDto) error
		Update(ctx context.Context, chartID uuid.UUID, dto dto.ChartUpdateDto) error
		Delete(ctx context.Context, chartID uuid.UUID) error
	}

	Dashboard interface {
		GetAll(ctx context.Context) ([]models.Dashboard, error)
		GetByTemplateID(ctx context.Context, templateID uuid.UUID) ([]models.Dashboard, error)
		GetByID(ctx context.Context, visualizationID uuid.UUID) (models.Dashboard, error)
		GetByShareID(ctx context.Context, shareID uuid.UUID) (models.Dashboard, error)
		Create(ctx context.Context, dto dto.VisualizationCreateDto) (uuid.UUID, error)
		Update(ctx context.Context, visualizationID uuid.UUID, dto dto.VisualizationUpdateDto) error
		IncrementViewCount(ctx context.Context, visualizationID uuid.UUID) error
		Delete(ctx context.Context, visualizationID uuid.UUID) error
	}

	Deps struct {
		Repo *repository.Repository
	}

	Service struct {
		User
		Template
		Canvas
		Chart
		Dashboard
	}
)

func New(log *slog.Logger, deps Deps) *Service {
	return &Service{
		User:      NewUserService(log, deps.Repo.User),
		Template:  NewTemplateService(log, deps.Repo.Template),
		Canvas:    NewCanvasService(log, deps.Repo.Canvas),
		Chart:     NewChartService(log, deps.Repo.Chart),
		Dashboard: NewVisualizationService(log, deps.Repo.Dashboard),
	}
}
