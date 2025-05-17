package usecase

import (
	"context"
	"log/slog"
	"visualizer-go/internal/domains/dashboards"

	"github.com/google/uuid"
)

type dashboardUsecase struct {
	log  *slog.Logger
	repo dashboards.Repository
}

func NewVisualizationService(log *slog.Logger, repo dashboards.Repository) dashboards.Usecase {
	return &dashboardUsecase{
		log:  log,
		repo: repo,
	}
}

func (vs *dashboardUsecase) GetAll(ctx context.Context) ([]*dashboards.Dashboard, error) {
	return vs.repo.GetAll(ctx)
}
func (vs *dashboardUsecase) GetByTemplateID(ctx context.Context, templateID uuid.UUID) ([]*dashboards.Dashboard, error) {
	return vs.repo.GetByTemplateID(ctx, templateID)
}

func (vs *dashboardUsecase) GetByID(ctx context.Context, visualizationID uuid.UUID) (*dashboards.Dashboard, error) {
	return vs.repo.GetByID(ctx, visualizationID)
}

func (vs *dashboardUsecase) GetByShareID(ctx context.Context, shareID uuid.UUID) (*dashboards.Dashboard, error) {
	return vs.repo.GetByShareID(ctx, shareID)
}

func (vs *dashboardUsecase) Create(ctx context.Context, dto dashboards.VisualizationCreateDto) (uuid.UUID, error) {
	return vs.repo.Create(ctx, dto)
}
func (vs *dashboardUsecase) Update(ctx context.Context, visualizationID uuid.UUID, dto dashboards.VisualizationUpdateDto) error {
	return vs.repo.Update(ctx, visualizationID, dto)
}

func (vs *dashboardUsecase) IncrementViewCount(ctx context.Context, visualizationID uuid.UUID) error {
	return vs.repo.IncrementViewCount(ctx, visualizationID)
}

func (vs *dashboardUsecase) Delete(ctx context.Context, visualizationID uuid.UUID) error {
	return vs.repo.Delete(ctx, visualizationID)
}
