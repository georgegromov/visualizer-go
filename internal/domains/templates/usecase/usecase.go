package usecase

import (
	"context"
	"log/slog"
	"visualizer-go/internal/domains/templates"

	"github.com/google/uuid"
)

type templateUsecase struct {
	log  *slog.Logger
	repo templates.Repository
}

func NewTemplateService(log *slog.Logger, repo templates.Repository) templates.Usecase {
	return &templateUsecase{
		log:  log,
		repo: repo,
	}
}

func (ts *templateUsecase) GetAll(ctx context.Context) ([]*templates.Template, error) {
	// const op = "service.TemplateService.GetAll"
	return ts.repo.GetAll(ctx)
}
func (ts *templateUsecase) GetByID(ctx context.Context, templateID uuid.UUID) (*templates.Template, error) {
	// const op = "service.TemplateService.GetByID"
	return ts.repo.GetByID(ctx, templateID)
}
func (ts *templateUsecase) Create(ctx context.Context, template *templates.Template) (uuid.UUID, error) {
	// const op = "service.TemplateService.Create"
	return ts.repo.Create(ctx, template)
}
func (ts *templateUsecase) Update(ctx context.Context, templateID uuid.UUID, dto *templates.TemplateUpdateDto) error {
	// const op = "service.TemplateService.Update"
	return ts.repo.Update(ctx, templateID, dto)
}
