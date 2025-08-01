package service

import (
	"context"
	"log/slog"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/models"
	"visualizer-go/internal/repository"

	"github.com/google/uuid"
)

type TemplateService struct {
	log  *slog.Logger
	repo repository.Template
}

func NewTemplateService(log *slog.Logger, repo repository.Template) *TemplateService {
	return &TemplateService{
		log:  log,
		repo: repo,
	}
}

func (ts *TemplateService) GetAll(ctx context.Context, withCanvases bool) ([]models.Template, error) {
	const op = "service.TemplateService.GetAll"
	return ts.repo.GetAll(ctx, withCanvases)
}
func (ts *TemplateService) GetByID(ctx context.Context, templateID uuid.UUID) (models.Template, error) {
	const op = "service.TemplateService.GetByID"
	return ts.repo.GetByID(ctx, templateID)
}
func (ts *TemplateService) Create(ctx context.Context, dto dto.TemplateCreateDto) (uuid.UUID, error) {
	const op = "service.TemplateService.Create"
	return ts.repo.Create(ctx, dto)
}
func (ts *TemplateService) Update(ctx context.Context, templateID uuid.UUID, dto dto.TemplateUpdateDto) error {
	const op = "service.TemplateService.Update"
	return ts.repo.Update(ctx, templateID, dto)
}
