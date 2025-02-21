package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/models"
)

type (
	Template interface {
		GetAll(ctx context.Context) ([]models.Template, error)
		GetByID(ctx context.Context, templateID uuid.UUID) (models.Template, error)
		Create(ctx context.Context, dto dto.TemplateCreateDto) error
		Update(ctx context.Context, templateID uuid.UUID, dto dto.TemplateUpdateDto) error
	}

	User interface {
		GetByID(ctx context.Context, userID uuid.UUID) (models.User, error)
		GetByUsername(ctx context.Context, username string) (models.User, error)
		Create(ctx context.Context, dto dto.UserCreateDto) error
		Update(ctx context.Context, userID uuid.UUID, dto dto.UserUpdateDto) error
	}

	Repository struct {
		Template
		User
	}
)

func New(log *slog.Logger, db *sqlx.DB) *Repository {
	return &Repository{
		Template: NewTemplateRepo(log, db),
		User:     NewUserRepo(log, db),
	}
}
