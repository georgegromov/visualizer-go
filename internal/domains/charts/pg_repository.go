package charts

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetByID(ctx context.Context, chartID uuid.UUID) (*Chart, error)
	GetByCanvasID(ctx context.Context, canvasID uuid.UUID) ([]*Chart, error)
	Create(ctx context.Context, dto ChartCreateDto) error
	Update(ctx context.Context, chartID uuid.UUID, dto ChartUpdateDto) error
	Delete(ctx context.Context, chartID uuid.UUID) error
}
