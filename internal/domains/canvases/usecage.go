package canvases

import (
	"context"

	"github.com/google/uuid"
)

type Usecase interface {
	Create(ctx context.Context, dto CanvasCreateDto) error
	GetCanvasesByTemplateID(ctx context.Context, templateID uuid.UUID) ([]*Canvas, error)
	Update(ctx context.Context, canvasID uuid.UUID, dto CanvasUpdateDto) error
	Delete(ctx context.Context, canvasID uuid.UUID) error
}
