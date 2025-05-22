package templates

import (
	"time"
	"visualizer-go/internal/domains/measurements"

	"github.com/google/uuid"
)

type Template struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Description *string    `json:"description" db:"description"`
	IsDeleted   bool       `json:"isDeleted" db:"is_deleted"`
	CreatorID   uuid.UUID  `json:"creatorId" db:"creator_id" validate:"required"`
	Uses        *uint      `json:"uses" db:"uses"`
	UpdatedAt   *time.Time `json:"updatedAt" db:"updated_at"`
	CreatedAt   time.Time  `json:"createdAt" db:"created_at"`
}

type AnalysisCanvas struct {
	ID     uuid.UUID       `json:"id" validate:"required"`
	Charts []AnalysisChart `json:"charts" validate:"required,dive"`
}

type AnalysisChart struct {
	ID           uuid.UUID             `json:"id" validate:"required"`
	Type         string                `json:"type" validate:"required"`
	Measurements []AnalysisMeasurement `json:"measurements" validate:"required, dive"`
	CanvasID     uuid.UUID             `json:"canvasId" validate:"required"`
}

type AnalysisMeasurement struct {
	ID        uuid.UUID            `json:"id" validate:"required"`
	Content   measurements.Content `json:"content" validate:"required"`
	CreatedAt string               `json:"createdAt" validate:"required"`
	ChartID   uuid.UUID            `json:"chartId" validate:"required"`
}
