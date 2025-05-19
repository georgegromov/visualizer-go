package measurements

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx/types"
)

type MeasurementCreateDTO struct {
	ChartID uuid.UUID      `json:"chartId" db:"chart_id" validate:"required"`
	Content types.JSONText `json:"content" db:"content" validate:"required"`
}

type MeasurementUpdateDTO struct {
	Content types.JSONText `json:"content" db:"content" validate:"required"`
}
