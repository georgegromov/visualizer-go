package charts

import "github.com/google/uuid"

type ChartCreateDto struct {
	Type     string    `json:"type" db:"type"`
	CanvasID uuid.UUID `json:"canvasId" db:"canvas_id"`
}

type ChartUpdateDto struct {
	Name         *string      `json:"name" db:"name"`
	Measurements *interface{} `json:"measurements" db:"measurements"`
}
