package charts

import (
	"time"

	"github.com/google/uuid"
)

type Chart struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	Name      *string    `json:"name" db:"name"`
	Type      string     `json:"type" db:"type"`
	CanvasID  *uuid.UUID `json:"canvasId" db:"canvas_id"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
}
