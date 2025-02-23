package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/lib/response"
)

var (
	ErrTemplateIDMissing          = errors.New("template ID is missing")
	ErrInvalidTemplateID          = errors.New("invalid template ID format")
	ErrFailedToFetchTemplates     = errors.New("failed to fetch templates")
	ErrTemplateNotFound           = errors.New("template not found")
	ErrFailedToCreateTemplate     = errors.New("failed to create template")
	ErrTemplateInvalidRequestData = errors.New("invalid template request data")
	ErrFailedToUpdateTemplate     = errors.New("failed to update template")
)

func (h *Handler) getAllTemplates(c *gin.Context) {
	const op = "handler.Handler.GetAllTemplatesHandler"

	templates, err := h.services.Template.GetAll(c.Request.Context())
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToFetchTemplates.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Templates fetched successfully", templates)
}

func (h *Handler) getTemplateByID(c *gin.Context) {
	const op = "handler.Handler.GetTemplateByIDHandler"

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

	template, err := h.services.Template.GetByID(c.Request.Context(), templateID)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusNotFound, ErrTemplateNotFound.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Template fetched successfully", template)
}

func (h *Handler) createTemplate(c *gin.Context) {
	const op = "handler.Handler.CreateTemplateHandler"

	var templateCreateDto dto.TemplateCreateDto
	if err := c.ShouldBindJSON(&templateCreateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrTemplateInvalidRequestData.Error(), err)
		return
	}

	templateID, err := h.services.Template.Create(c.Request.Context(), templateCreateDto)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToCreateTemplate.Error(), err)
		return
	}

	response.Success(c, http.StatusCreated, "Template created successfully", templateID)
}

func (h *Handler) updateTemplate(c *gin.Context) {
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

	if err = h.services.Template.Update(c.Request.Context(), templateID, templateUpdateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToUpdateTemplate.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Template updated successfully", nil)
}
