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
	ErrFailedToFetchCanvases = errors.New("failed to fetch canvases")
)

func (h *Handler) createCanvasHandler(c *gin.Context) {
	const op = "handler.Handler.createCanvasHandler"

	// TODO: дописать и исправить ошибки

	var canvasCreateDto dto.CanvasCreateDto
	if err := c.ShouldBindJSON(&canvasCreateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrVisualizationInvalidRequestData.Error(), err)
		return
	}

	err := h.services.Canvas.Create(c.Request.Context(), canvasCreateDto)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToCreateVisualization.Error(), err)
		return
	}

	response.Success(c, http.StatusCreated, "Canvas created successfully", gin.H{})
}

func (h *Handler) getCanvasByTemplateIdHandler(c *gin.Context) {

	const op = "handler.Handler.getCanvasByTemplateIdHandler"

	templateIDStr := c.Query("templateId")
	if templateIDStr == "" {
		h.log.Error(fmt.Sprintf("%s: %v", op, ErrTemplateIDMissing))
		response.Error(c, http.StatusBadRequest, ErrTemplateIDMissing.Error(), nil)
		return
	}

	templateID, err := uuid.Parse(templateIDStr)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrInvalidVisualizationID.Error(), err)
		return
	}

	canvases, err := h.services.Canvas.GetCanvasesByTemplateID(c.Request.Context(), templateID)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToFetchVisualizations.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Canvases fetched successfully", canvases)
}

func (h *Handler) updateCanvasHandler(c *gin.Context) {
	const op = "handler.Handler.updateCanvasHandler"

	canvasIDStr := c.Param("id")
	if canvasIDStr == "" {
		h.log.Error(fmt.Sprintf("%s: %v", op, ErrVisualizationIDMissing))
		response.Error(c, http.StatusBadRequest, ErrVisualizationIDMissing.Error(), nil)
		return
	}

	canvasID, err := uuid.Parse(canvasIDStr)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrInvalidVisualizationID.Error(), err)
		return
	}

	var canvasUpdateDto dto.CanvasUpdateDto
	if err = c.ShouldBindJSON(&canvasUpdateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrVisualizationInvalidRequestData.Error(), err)
		return
	}

	if err = h.services.Canvas.Update(c.Request.Context(), canvasID, canvasUpdateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToUpdateVisualization.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Canvas updated successfully", nil)
}

func (h *Handler) deleteCanvasHandler(c *gin.Context) {
	const op = "handler.Handler.deleteCanvasHandler"

	canvasIDStr := c.Param("id")
	if canvasIDStr == "" {
		h.log.Error(fmt.Sprintf("%s: %v", op, ErrVisualizationIDMissing))
		response.Error(c, http.StatusBadRequest, ErrVisualizationIDMissing.Error(), nil)
		return
	}

	canvasID, err := uuid.Parse(canvasIDStr)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrInvalidVisualizationID.Error(), err)
		return
	}

	if err = h.services.Canvas.Delete(c.Request.Context(), canvasID); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToDeleteVisualization.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Canvas deleted successfully", nil)
}
