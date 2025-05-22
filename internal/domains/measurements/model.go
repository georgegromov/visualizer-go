package measurements

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx/types"
)

type Measurement struct {
	ID        uuid.UUID      `json:"id" db:"id"`
	Content   types.JSONText `json:"content" db:"content"`
	ChartID   uuid.UUID      `json:"chartId" db:"chart_id"`
	UpdatedAt *time.Time     `json:"updatedAt" db:"updated_at"`
	CreatedAt time.Time      `json:"createdAt" db:"created_at"`
}

type Content struct {
	Connection Connection `json:"connection" validate:"required"`
	Series     Series     `json:"series" validate:"required"`
}

type Connection struct {
	ProjectID              uuid.UUID `json:"projectId" validate:"required"`
	ProjectName            string    `json:"projectName" validate:"required"`
	PlantID                uuid.UUID `json:"plantId" validate:"required"`
	PlantName              string    `json:"plantName" validate:"required"`
	AssetID                uuid.UUID `json:"assetId" validate:"required"`
	AssetName              string    `json:"assetName" validate:"required"`
	MeasurementID          uuid.UUID `json:"measurementId" validate:"required"`
	MeasurementName        string    `json:"measurementName" validate:"required"`
	MeasurementVersionID   uuid.UUID `json:"measurementVersionId" validate:"required"`
	MeasurementVersionName string    `json:"measurementVersionName" validate:"required"`
}

type Series struct {
	Name string `json:"name" validate:"required"`
	Type string `json:"type" validate:"required"`
	Step bool   `json:"step" validate:"required"`
	Dash bool   `json:"dash" validate:"required"`
}
