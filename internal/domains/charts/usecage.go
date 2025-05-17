package charts

import (
	"context"

	"github.com/google/uuid"
)

type Usecase interface {
	GetByCanvasID(ctx context.Context, canvasID uuid.UUID) ([]*Chart, error)
	Create(ctx context.Context, dto ChartCreateDto) error
	Update(ctx context.Context, chartID uuid.UUID, dto ChartUpdateDto) error
	Delete(ctx context.Context, chartID uuid.UUID) error
}
