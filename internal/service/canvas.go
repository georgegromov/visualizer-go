package service

import (
	"context"
	"log/slog"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/models"
	"visualizer-go/internal/repository"

	"github.com/google/uuid"
)

type CanvasService struct {
	log  *slog.Logger
	repo repository.Canvas
}

func NewCanvasService(log *slog.Logger, repo repository.Canvas) *CanvasService {
	return &CanvasService{log: log, repo: repo}
}

func (c *CanvasService) Create(ctx context.Context, dto dto.CanvasCreateDto) error {
	return c.repo.Create(ctx, dto)
}

func (c *CanvasService) GetCanvasesByTemplateID(ctx context.Context, templateID uuid.UUID) ([]*models.Canvas, error) {
	return c.repo.GetCanvasesByTemplateID(ctx, templateID)
}

func (c *CanvasService) Update(ctx context.Context, canvasID uuid.UUID, dto dto.CanvasUpdateDto) error {
	return c.repo.Update(ctx, canvasID, dto)
}

func (c *CanvasService) Delete(ctx context.Context, canvasID uuid.UUID) error {
	return c.repo.Delete(ctx, canvasID)
}
