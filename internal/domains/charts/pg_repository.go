package charts

import (
	"context"
	"visualizer-go/internal/dto"

	"github.com/google/uuid"
)

type Repository interface {
	GetByCanvasID(ctx context.Context, canvasID uuid.UUID) ([]*Chart, error)
	Create(ctx context.Context, dto dto.ChartCreateDto) error
	Update(ctx context.Context, chartID uuid.UUID, dto dto.ChartUpdateDto) error
	Delete(ctx context.Context, chartID uuid.UUID) error
}
