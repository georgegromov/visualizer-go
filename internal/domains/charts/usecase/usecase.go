package usecase

import (
	"context"
	"log/slog"
	"visualizer-go/internal/domains/charts"

	"github.com/google/uuid"
)

type chartUsecase struct {
	log  *slog.Logger
	repo charts.Repository
}

func NewChartService(log *slog.Logger, repo charts.Repository) charts.Usecase {
	return &chartUsecase{log: log, repo: repo}
}

func (c *chartUsecase) GetByID(ctx context.Context, chartID uuid.UUID) (*charts.Chart, error) {
	return c.repo.GetByID(ctx, chartID)
}

func (c *chartUsecase) GetByCanvasID(ctx context.Context, canvasID uuid.UUID) ([]*charts.Chart, error) {
	return c.repo.GetByCanvasID(ctx, canvasID)
}
func (c *chartUsecase) Create(ctx context.Context, dto charts.ChartCreateDto) error {
	return c.repo.Create(ctx, dto)
}
func (c *chartUsecase) Update(ctx context.Context, chartID uuid.UUID, dto charts.ChartUpdateDto) error {
	return c.repo.Update(ctx, chartID, dto)
}
func (c *chartUsecase) Delete(ctx context.Context, chartID uuid.UUID) error {
	return c.repo.Delete(ctx, chartID)
}
