package templates

import (
	"context"
	"visualizer-go/internal/dto"

	"github.com/google/uuid"
)

type Usecase interface {
	GetAll(ctx context.Context) ([]*Template, error)
	GetByID(ctx context.Context, templateID uuid.UUID) (*Template, error)
	Create(ctx context.Context, template *Template) (uuid.UUID, error)
	Update(ctx context.Context, templateID uuid.UUID, dto dto.TemplateUpdateDto) error
}
