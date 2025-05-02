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
	Template interface {
		GetAll(ctx context.Context, withCanvases bool) ([]models.Template, error)
		GetByID(ctx context.Context, templateID uuid.UUID) (models.Template, error)
		Create(ctx context.Context, dto dto.TemplateCreateDto) (uuid.UUID, error)
		Update(ctx context.Context, templateID uuid.UUID, dto dto.TemplateUpdateDto) error
	}

	User interface {
		GetByID(ctx context.Context, userID uuid.UUID) (models.User, error)
		GetByUsername(ctx context.Context, username string) (models.User, error)
		Create(ctx context.Context, dto dto.UserCreateDto) error
		Update(ctx context.Context, userID uuid.UUID, dto dto.UserUpdateDto) error
	}

	Visualization interface {
		GetAll(ctx context.Context) ([]models.Visualization, error)
		GetByTemplateID(ctx context.Context, templateID uuid.UUID) ([]models.Visualization, error)
		GetByID(ctx context.Context, visualizationID uuid.UUID) (models.Visualization, error)
		GetByShareID(ctx context.Context, shareID uuid.UUID) (models.Visualization, error)
		Create(ctx context.Context, dto dto.VisualizationCreateDto) (uuid.UUID, error)
		Update(ctx context.Context, visualizationID uuid.UUID, dto dto.VisualizationUpdateDto) error
    IncrementViewCount(ctx context.Context, visualizationID uuid.UUID) error
		Delete(ctx context.Context, visualizationID uuid.UUID) error
	}

	Repository struct {
		Template
		User
		Visualization
	}
)

func New(log *slog.Logger, db *sqlx.DB) *Repository {
	return &Repository{
		Template:      NewTemplateRepo(log, db),
		User:          NewUserRepo(log, db),
		Visualization: NewVisualizationRepo(log, db),
	}
}
