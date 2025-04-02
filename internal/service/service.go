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
	Template interface {
		GetAll(ctx context.Context, withCanvases bool) ([]models.Template, error)
		GetByID(ctx context.Context, templateID uuid.UUID) (models.Template, error)
		Create(ctx context.Context, dto dto.TemplateCreateDto) (uuid.UUID, error)
		Update(ctx context.Context, templateID uuid.UUID, dto dto.TemplateUpdateDto) error
	}

	User interface {
		Login(ctx context.Context, dto dto.UserLoginDto) (models.User, string, error)
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
		Delete(ctx context.Context, visualizationID uuid.UUID) error
	}

	Deps struct {
		Repo *repository.Repository
	}

	Service struct {
		Template
		User
		Visualization
	}
)

func New(log *slog.Logger, deps Deps) *Service {
	return &Service{
		Template:      NewTemplateService(log, deps.Repo.Template),
		User:          NewUserService(log, deps.Repo.User),
		Visualization: NewVisualizationService(log, deps.Repo.Visualization),
	}
}
