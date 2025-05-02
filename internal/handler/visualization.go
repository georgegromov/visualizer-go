package handler

import (
	"errors"
	"fmt"
	"net/http"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/lib/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	ErrVisualizationIDMissing                  = errors.New("visualization ID is missing")
	ErrInvalidVisualizationID                  = errors.New("invalid visualization ID format")
	ErrFailedToFetchVisualizations             = errors.New("failed to fetch visualization")
	ErrVisualizationNotFound                   = errors.New("visualization not found")
	ErrFailedToCreateVisualization             = errors.New("failed to create visualization")
	ErrVisualizationInvalidRequestData         = errors.New("invalid visualization request data")
	ErrFailedToUpdateVisualization             = errors.New("failed to update visualization")
	ErrFailedToDeleteVisualization             = errors.New("failed to delete visualization")
	ErrFailedToIncrementViewCountVisualization = errors.New("failed to increment view count visualization")
)

// TODO: rename template -> visualization

func (h *Handler) getAllVisualizations(c *gin.Context) {
	const op = "handler.Handler.getAllVisualizations"

	templates, err := h.services.Visualization.GetAll(c.Request.Context())
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToFetchVisualizations.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Visualizations fetched successfully", templates)
}

func (h *Handler) getVisualizationsByTemplateID(c *gin.Context) {
	const op = "handler.Handler.getVisualizationsByTemplateID"

	templateIDStr := c.Param("id")
	if templateIDStr == "" {
		h.log.Error(fmt.Sprintf("%s: %v", op, ErrVisualizationIDMissing))
		response.Error(c, http.StatusBadRequest, ErrVisualizationIDMissing.Error(), nil)
		return
	}

	templateID, err := uuid.Parse(templateIDStr)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrInvalidVisualizationID.Error(), err)
		return
	}

	templates, err := h.services.Visualization.GetByTemplateID(c.Request.Context(), templateID)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToFetchVisualizations.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Visualizations fetched successfully", templates)
}

func (h *Handler) getVisualizationByID(c *gin.Context) {
	const op = "handler.Handler.getVisualizationByID"

	templateIDStr := c.Param("id")
	if templateIDStr == "" {
		h.log.Error(fmt.Sprintf("%s: %v", op, ErrVisualizationIDMissing))
		response.Error(c, http.StatusBadRequest, ErrVisualizationIDMissing.Error(), nil)
		return
	}

	templateID, err := uuid.Parse(templateIDStr)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrInvalidVisualizationID.Error(), err)
		return
	}

	template, err := h.services.Visualization.GetByID(c.Request.Context(), templateID)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusNotFound, ErrVisualizationNotFound.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Visualization fetched successfully", template)
}

func (h *Handler) getVisualizationByShareID(c *gin.Context) {
	const op = "handler.Handler.getVisualizationByShareID"

	templateIDStr := c.Param("id")
	if templateIDStr == "" {
		h.log.Error(fmt.Sprintf("%s: %v", op, ErrVisualizationIDMissing))
		response.Error(c, http.StatusBadRequest, ErrVisualizationIDMissing.Error(), nil)
		return
	}

	templateID, err := uuid.Parse(templateIDStr)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrInvalidVisualizationID.Error(), err)
		return
	}

	template, err := h.services.Visualization.GetByShareID(c.Request.Context(), templateID)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusNotFound, ErrVisualizationNotFound.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Visualization fetched successfully", template)
}

func (h *Handler) createVisualization(c *gin.Context) {
	const op = "handler.Handler.createVisualization"

	var visualizationCreateDto dto.VisualizationCreateDto
	if err := c.ShouldBindJSON(&visualizationCreateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrVisualizationInvalidRequestData.Error(), err)
		return
	}

	templateID, err := h.services.Visualization.Create(c.Request.Context(), visualizationCreateDto)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToCreateVisualization.Error(), err)
		return
	}

	response.Success(c, http.StatusCreated, "Visualization created successfully", templateID)
}

func (h *Handler) updateVisualization(c *gin.Context) {
	const op = "handler.Handler.updateVisualization"

	templateIDStr := c.Param("id")
	if templateIDStr == "" {
		h.log.Error(fmt.Sprintf("%s: %v", op, ErrVisualizationIDMissing))
		response.Error(c, http.StatusBadRequest, ErrVisualizationIDMissing.Error(), nil)
		return
	}

	templateID, err := uuid.Parse(templateIDStr)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrInvalidVisualizationID.Error(), err)
		return
	}

	var visualizationUpdateDto dto.VisualizationUpdateDto
	if err = c.ShouldBindJSON(&visualizationUpdateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrVisualizationInvalidRequestData.Error(), err)
		return
	}

	if err = h.services.Visualization.Update(c.Request.Context(), templateID, visualizationUpdateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToUpdateVisualization.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Visualization updated successfully", nil)
}

func (h *Handler) metric(c *gin.Context) {
	const op = "handler.Handler.metric"

	templateIDStr := c.Param("id")
	if templateIDStr == "" {
		h.log.Error(fmt.Sprintf("%s: %v", op, ErrVisualizationIDMissing))
		response.Error(c, http.StatusBadRequest, ErrVisualizationIDMissing.Error(), nil)
		return
	}

	templateID, err := uuid.Parse(templateIDStr)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrInvalidVisualizationID.Error(), err)
		return
	}

	if err = h.services.IncrementViewCount(c.Request.Context(), templateID); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToIncrementViewCountVisualization.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "view count incremented", nil)
}

func (h *Handler) deleteVisualization(c *gin.Context) {
	const op = "handler.Handler.deleteVisualization"

	templateIDStr := c.Param("id")
	if templateIDStr == "" {
		h.log.Error(fmt.Sprintf("%s: %v", op, ErrVisualizationIDMissing))
		response.Error(c, http.StatusBadRequest, ErrVisualizationIDMissing.Error(), nil)
		return
	}

	templateID, err := uuid.Parse(templateIDStr)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrInvalidVisualizationID.Error(), err)
		return
	}

	if err = h.services.Visualization.Delete(c.Request.Context(), templateID); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToDeleteVisualization.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Visualization deleted successfully", nil)
}
