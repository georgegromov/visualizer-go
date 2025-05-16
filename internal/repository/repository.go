package repository

import (
	"context"
	"log/slog"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type (
	User interface {
		GetByID(ctx context.Context, userID uuid.UUID) (models.User, error)
		GetByUsername(ctx context.Context, username string) (*models.User, error)
		Create(ctx context.Context, user *models.User) error
		Update(ctx context.Context, userID uuid.UUID, dto dto.UserUpdateDto) error
	}
	Template interface {
		GetAll(ctx context.Context) ([]models.Template, error)
		GetByID(ctx context.Context, templateID uuid.UUID) (models.Template, error)
		Create(ctx context.Context, dto dto.TemplateCreateDto) (uuid.UUID, error)
		Update(ctx context.Context, templateID uuid.UUID, dto dto.TemplateUpdateDto) error
	}

	Canvas interface {
		GetCanvasesByTemplateID(ctx context.Context, templateID uuid.UUID) ([]models.Canvas, error)
		Create(ctx context.Context, dto dto.CanvasCreateDto) error
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

	Repository struct {
		User
		Template
		Canvas
		Chart
		Dashboard
	}
)

func New(log *slog.Logger, db *sqlx.DB) *Repository {
	return &Repository{
		User:      NewUserRepo(log, db),
		Template:  NewTemplateRepo(log, db),
		Canvas:    NewCanvasRepo(log, db),
		Chart:     NewChartRepo(log, db),
		Dashboard: NewVisualizationRepo(log, db),
	}
}
