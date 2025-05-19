package usecase

import (
	"context"
	"log/slog"
	"visualizer-go/internal/domains/measurements"

	"github.com/google/uuid"
)

type measurementUsecase struct {
	log  *slog.Logger
	repo measurements.Repository
}

func NewMeasurementUsecase(log *slog.Logger, repo measurements.Repository) measurements.Usecase {
	return &measurementUsecase{log: log, repo: repo}
}

func (uc *measurementUsecase) GetByChartID(ctx context.Context, chartID uuid.UUID) ([]*measurements.Measurement, error) {
	return uc.repo.GetByChartID(ctx, chartID)
}
func (uc *measurementUsecase) Create(ctx context.Context, dto *measurements.MeasurementCreateDTO) error {
	return uc.repo.Create(ctx, dto)
}
func (uc *measurementUsecase) Update(ctx context.Context, measurementID uuid.UUID, dto *measurements.MeasurementUpdateDTO) error {
	return uc.repo.Update(ctx, measurementID, dto)
}
func (uc *measurementUsecase) Delete(ctx context.Context, measurementID uuid.UUID) error {
	return uc.repo.Delete(ctx, measurementID)
}
