package measurements

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetByChartID(ctx context.Context, chartID uuid.UUID) ([]*Measurement, error)
	Create(ctx context.Context, dto *MeasurementCreateDTO) error
	Update(ctx context.Context, measurementID uuid.UUID, dto *MeasurementUpdateDTO) error
	Delete(ctx context.Context, measurementID uuid.UUID) error
}
