package service

import (
	"context"
	"log/slog"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/models"
	"visualizer-go/internal/repository"

	"github.com/google/uuid"
)

type VisualizationService struct {
	log  *slog.Logger
	repo repository.Visualization
}

func NewVisualizationService(log *slog.Logger, repo repository.Visualization) *VisualizationService {
	return &VisualizationService{
		log:  log,
		repo: repo,
	}
}

func (vs *VisualizationService) GetAll(ctx context.Context) ([]models.Visualization, error) {
	const op = "service.VisualizationService.GetAll"
	return vs.repo.GetAll(ctx)
}
func (vs *VisualizationService) GetByTemplateID(ctx context.Context, templateID uuid.UUID) ([]models.Visualization, error) {
	const op = "service.VisualizationService.GetByTemplateID"
	return vs.repo.GetByTemplateID(ctx, templateID)
}

func (vs *VisualizationService) GetByID(ctx context.Context, visualizationID uuid.UUID) (models.Visualization, error) {
	const op = "service.VisualizationService.GetByID"
	return vs.repo.GetByID(ctx, visualizationID)
}

func (vs *VisualizationService) GetByShareID(ctx context.Context, shareID uuid.UUID) (models.Visualization, error) {
	const op = "service.VisualizationService.GetByShareID"
	return vs.repo.GetByShareID(ctx, shareID)
}

func (vs *VisualizationService) Create(ctx context.Context, dto dto.VisualizationCreateDto) (uuid.UUID, error) {
	const op = "service.VisualizationService.Create"
	return vs.repo.Create(ctx, dto)
}
func (vs *VisualizationService) Update(ctx context.Context, visualizationID uuid.UUID, dto dto.VisualizationUpdateDto) error {
	const op = "service.VisualizationService.Update"
	return vs.repo.Update(ctx, visualizationID, dto)
}

func (vs *VisualizationService) Delete(ctx context.Context, visualizationID uuid.UUID) error {
	const op = "service.VisualizationService.Delete"
	return vs.repo.Delete(ctx, visualizationID)
}
