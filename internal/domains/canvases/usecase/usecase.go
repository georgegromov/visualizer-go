package usecase

import (
	"context"
	"log/slog"
	"visualizer-go/internal/domains/canvases"

	"github.com/google/uuid"
)

type canvasUsecase struct {
	log  *slog.Logger
	repo canvases.Repository
}

func NewCanvasUsecase(log *slog.Logger, repo canvases.Repository) canvases.Usecase {
	return &canvasUsecase{log: log, repo: repo}
}

func (c *canvasUsecase) Create(ctx context.Context, dto canvases.CanvasCreateDto) error {
	return c.repo.Create(ctx, dto)
}

func (c *canvasUsecase) GetCanvasesByTemplateID(ctx context.Context, templateID uuid.UUID) ([]*canvases.Canvas, error) {
	return c.repo.GetCanvasesByTemplateID(ctx, templateID)
}

func (c *canvasUsecase) Update(ctx context.Context, canvasID uuid.UUID, dto canvases.CanvasUpdateDto) error {
	return c.repo.Update(ctx, canvasID, dto)
}

func (c *canvasUsecase) Delete(ctx context.Context, canvasID uuid.UUID) error {
	return c.repo.Delete(ctx, canvasID)
}
