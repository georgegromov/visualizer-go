package dashboards

import (
	"context"
	"visualizer-go/internal/dto"

	"github.com/google/uuid"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*Dashboard, error)
	GetByTemplateID(ctx context.Context, templateID uuid.UUID) ([]*Dashboard, error)
	GetByID(ctx context.Context, visualizationID uuid.UUID) (*Dashboard, error)
	GetByShareID(ctx context.Context, shareID uuid.UUID) (*Dashboard, error)
	Create(ctx context.Context, dto dto.VisualizationCreateDto) (uuid.UUID, error)
	Update(ctx context.Context, visualizationID uuid.UUID, dto dto.VisualizationUpdateDto) error
	IncrementViewCount(ctx context.Context, visualizationID uuid.UUID) error
	Delete(ctx context.Context, visualizationID uuid.UUID) error
}
