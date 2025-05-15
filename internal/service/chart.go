package service

import (
	"context"
	"log/slog"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/models"
	"visualizer-go/internal/repository"

	"github.com/google/uuid"
)

type ChartService struct {
	log  *slog.Logger
	repo repository.Chart
}

func NewChartService(log *slog.Logger, repo repository.Chart) *ChartService {
	return &ChartService{log: log, repo: repo}
}

func (c *ChartService) GetByCanvasID(ctx context.Context, canvasID uuid.UUID) ([]models.Chart, error) {
	return c.repo.GetByCanvasID(ctx, canvasID)
}
func (c *ChartService) Create(ctx context.Context, dto dto.ChartCreateDto) error {
	return c.repo.Create(ctx, dto)
}
func (c *ChartService) Update(ctx context.Context, chartID uuid.UUID, dto dto.ChartUpdateDto) error {
	return c.repo.Update(ctx, chartID, dto)
}
func (c *ChartService) Delete(ctx context.Context, chartID uuid.UUID) error {
	return c.repo.Delete(ctx, chartID)
}
