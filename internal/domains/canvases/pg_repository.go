package canvases

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetCanvasByID(ctx context.Context, canvasID uuid.UUID) (*Canvas, error)
	GetCanvasesByTemplateID(ctx context.Context, templateID uuid.UUID) ([]*Canvas, error)
	Create(ctx context.Context, dto CanvasCreateDto) error
	Update(ctx context.Context, canvasID uuid.UUID, dto CanvasUpdateDto) error
	Delete(ctx context.Context, canvasID uuid.UUID) error
}
