package canvases

import (
	"context"
	"visualizer-go/internal/dto"

	"github.com/google/uuid"
)

type Repository interface {
	GetCanvasesByTemplateID(ctx context.Context, templateID uuid.UUID) ([]*Canvas, error)
	Create(ctx context.Context, dto dto.CanvasCreateDto) error
	Update(ctx context.Context, canvasID uuid.UUID, dto dto.CanvasUpdateDto) error
	Delete(ctx context.Context, canvasID uuid.UUID) error
}
