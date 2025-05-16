package delivery

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"visualizer-go/internal/domains/templates"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/response"
	"visualizer-go/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type templateHandler struct {
	log        *slog.Logger
	templateUC templates.Usecase
}

func NewTemplateHandler(log *slog.Logger, templateUC templates.Usecase) templates.Handler {
	return &templateHandler{log: log, templateUC: templateUC}
}

var (
	ErrTemplateIDMissing          = errors.New("template ID is missing")
	ErrInvalidTemplateID          = errors.New("invalid template ID format")
	ErrFailedToFetchTemplates     = errors.New("failed to fetch templates")
	ErrTemplateNotFound           = errors.New("template not found")
	ErrFailedToCreateTemplate     = errors.New("failed to create template")
	ErrTemplateInvalidRequestData = errors.New("invalid template request data")
	ErrFailedToUpdateTemplate     = errors.New("failed to update template")
)

// getTemplates godoc
//
// @Summary Returns all templates
// @Description Returns all templates
// @Tags Templates
// @Accept json
// @Produce json
// @Success 200 {object} []models.Template
// @Router /templates [get]
func (h *templateHandler) HandleGet(c *gin.Context) {
	const op = "handler.Handler.GetAllTemplatesHandler"

	templates, err := h.templateUC.GetAll(c.Request.Context())
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToFetchTemplates.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Templates fetched successfully", templates)
}

// getTemplate godoc
//
// @Summary Returns template
// @Description Returns template
// @Tags Templates
// @Accept json
// @Produce json
// @Success 200 {object} models.Template
// @Router /templates/{id} [get]
func (h *templateHandler) HandleGetById(c *gin.Context) {
	const op = "handler.Handler.GetTemplateByIDHandler"

	templateIDStr := c.Param("templateId")
	if templateIDStr == "" {
		h.log.Error(fmt.Sprintf("%s: %v", op, ErrTemplateIDMissing))
		response.Error(c, http.StatusBadRequest, ErrTemplateIDMissing.Error(), nil)
		return
	}

	templateID, err := uuid.Parse(templateIDStr)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrInvalidTemplateID.Error(), err)
		return
	}

	template, err := h.templateUC.GetByID(c.Request.Context(), templateID)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusNotFound, ErrTemplateNotFound.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Template fetched successfully", template)
}

// createTemplate godoc
//
// @Summary Create template
// @Description Create template
// @Tags Templates
// @Accept json
// @Produce json
// @Success 201
// @Router /templates/{id} [post]
func (h *templateHandler) HandleCreate(ctx *gin.Context) {
	const op = "handler.Handler.CreateTemplateHandler"

	user, err := utils.GetUserFromCtx(ctx)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: no user set in context: %v", op, err))
		response.Error(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	template := &templates.Template{}
	template.CreatorID = user.ID
	if err := utils.ReadRequestBody(ctx, template); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	templateID, err := h.templateUC.Create(ctx.Request.Context(), template)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusInternalServerError, ErrFailedToCreateTemplate.Error(), err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Template created successfully", templateID)
}

// updateTemplate godoc
//
// @Summary Update template
// @Description Update template
// @Tags Templates
// @Accept json
// @Produce json
// @Success 200
// @Router /templates/{id} [patch]
func (h *templateHandler) HandleUpdate(c *gin.Context) {
	const op = "handler.Handler.UpdateTemplateHandler"

	templateIDStr := c.Param("id")
	if templateIDStr == "" {
		h.log.Error(fmt.Sprintf("%s: %v", op, ErrTemplateIDMissing))
		response.Error(c, http.StatusBadRequest, ErrTemplateIDMissing.Error(), nil)
		return
	}

	templateID, err := uuid.Parse(templateIDStr)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrInvalidTemplateID.Error(), err)
		return
	}

	var templateUpdateDto dto.TemplateUpdateDto
	if err = c.ShouldBindJSON(&templateUpdateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrTemplateInvalidRequestData.Error(), err)
		return
	}

	if err = h.templateUC.Update(c.Request.Context(), templateID, templateUpdateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToUpdateTemplate.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Template updated successfully", nil)
}
