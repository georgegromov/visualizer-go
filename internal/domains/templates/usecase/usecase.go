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

func (uc *templateUsecase) GetAll(ctx context.Context) ([]*templates.Template, error) {
	// const op = "service.TemplateService.GetAll"
	return uc.repo.GetAll(ctx)
}
func (uc *templateUsecase) GetByID(ctx context.Context, templateID uuid.UUID) (*templates.Template, error) {
	// const op = "service.TemplateService.GetByID"
	return uc.repo.GetByID(ctx, templateID)
}
func (uc *templateUsecase) Create(ctx context.Context, template *templates.Template) (uuid.UUID, error) {
	// const op = "service.TemplateService.Create"
	return uc.repo.Create(ctx, template)
}
func (uc *templateUsecase) Update(ctx context.Context, templateID uuid.UUID, dto *templates.TemplateUpdateDto) error {
	// const op = "service.TemplateService.Update"
	return uc.repo.Update(ctx, templateID, dto)
}

func (uc *templateUsecase) SaveAs(ctx context.Context, dto *templates.TemplateSaveAsDTO) error {

	// Создать шаблон, получить templateID
	// err := uc.Create(ctx)

	// for _, canvas := range dto.Canvases {

	// 	cnvs := &canvases.CanvasCreateDto{
	// 		TemplateID: uuid.New(),
	// 	}

	// 	// Создать канвас, получить CanvasID

	// 	for _, chart := range canvas.Charts {

	// 		chrt := &charts.ChartCreateDto{
	// 			Type:     chart.Type,
	// 			CanvasID: chart.CanvasID,
	// 		}

	// 		// Создать chart, получить ChartID

	// 		for _, measurement := range chart.Measurements {

	// 			contentJson, err := json.Marshal(measurement.Content)
	// 			if err != nil {
	// 				return fmt.Errorf("failed to marshal measurement content: %w", err)
	// 			}

	// 			msrmnt := &measurements.MeasurementCreateDTO{
	// 				ChartID: chart.ID,
	// 				Content: types.JSONText(contentJson),
	// 			}

	// 			// Создать measurement
	// 		}
	// 	}
	// }

	return nil
}
