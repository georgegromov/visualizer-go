package diagrams

import (
	"context"

	"github.com/google/uuid"
)

type Usecase interface {
	GetAll(ctx context.Context) ([]*Diagram, error)
	GetByID(ctx context.Context, diagramID uuid.UUID) (*Diagram, error)
	Create(ctx context.Context, dto *DiagramCreateDTO) (uuid.UUID, error)
	Update(ctx context.Context, diagramID uuid.UUID, dto *DiagramUpdateDTO) error
	Delete(ctx context.Context, diagramID uuid.UUID) error
}
