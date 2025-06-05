package usecase

import (
	"context"
	"log/slog"
	"visualizer-go/internal/domains/diagrams"

	"github.com/google/uuid"
)

type diagramUsecase struct {
	log  *slog.Logger
	repo diagrams.Repository
}

func NewDiagramUsecase(log *slog.Logger, repo diagrams.Repository) diagrams.Usecase {
	return &diagramUsecase{log: log, repo: repo}
}

func (uc *diagramUsecase) GetAll(ctx context.Context) ([]*diagrams.Diagram, error) {
	return uc.repo.GetAll(ctx)
}
func (uc *diagramUsecase) GetByID(ctx context.Context, diagramID uuid.UUID) (*diagrams.Diagram, error) {
	return uc.repo.GetByID(ctx, diagramID)
}
func (uc *diagramUsecase) Create(ctx context.Context, dto *diagrams.DiagramCreateDTO) (uuid.UUID, error) {
	return uc.repo.Create(ctx, dto)
}
func (uc *diagramUsecase) Update(ctx context.Context, diagramID uuid.UUID, dto *diagrams.DiagramUpdateDTO) error {
	return uc.repo.Update(ctx, diagramID, dto)
}
func (uc *diagramUsecase) Delete(ctx context.Context, diagramID uuid.UUID) error {
	return uc.repo.Delete(ctx, diagramID)
}
