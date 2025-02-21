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

	User interface {
		Login(ctx context.Context, dto dto.UserLoginDto) (models.User, string, error)
		GetByID(ctx context.Context, userID uuid.UUID) (models.User, error)
		GetByUsername(ctx context.Context, username string) (models.User, error)
		Create(ctx context.Context, dto dto.UserCreateDto) error
		Update(ctx context.Context, userID uuid.UUID, dto dto.UserUpdateDto) error
	}

	Deps struct {
		Repo *repository.Repository
	}

	Service struct {
		Template
		User
	}
)

func New(log *slog.Logger, deps Deps) *Service {
	return &Service{
		Template: NewTemplateService(log, deps.Repo.Template),
		User:     NewUserService(log, deps.Repo.User),
	}
}
