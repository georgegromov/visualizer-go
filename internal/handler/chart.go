package handler

import (
	"errors"
	"fmt"
	"net/http"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	ErrCanvasIdMissing     = errors.New("canvas id is missing")
	ErrChartIdMissing      = errors.New("chart id is missing")
	ErrInvalidChartId      = errors.New("invalid chart id")
	ErrFailedToFetchCharts = errors.New("failed to fetch charts")
	ErrFailedToCreateChart = errors.New("failed to create chart")
	ErrFailedToDeleteChart = errors.New("failed to delete chart")
	ErrBadRequest          = errors.New("bad request")
)

func (h *Handler) getChartsByCanvasIdHanlder(c *gin.Context) {
	const op = "handler.Handler.getChartsByCanvasIdHanlder"

	canvasIdStr := c.Query("canvasId")
	if canvasIdStr == "" {
		h.log.Error(fmt.Sprintf("%s: %v", op, ErrCanvasIdMissing))
		response.Error(c, http.StatusBadRequest, ErrCanvasIdMissing.Error(), nil)
		return
	}

	canvasId, err := uuid.Parse(canvasIdStr)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrInvalidVisualizationID.Error(), err)
		return
	}

	charts, err := h.services.Chart.GetByCanvasID(c.Request.Context(), canvasId)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToFetchCharts.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Charts fetched successfully", charts)
}

func (h *Handler) createChartHanlder(c *gin.Context) {
	const op = "handler.Handler.createChartHanlder"

	var chartCreateDto dto.ChartCreateDto
	if err := c.ShouldBindJSON(&chartCreateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrBadRequest.Error(), err)
		return
	}

	err := h.services.Chart.Create(c.Request.Context(), chartCreateDto)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToCreateChart.Error(), err)
		return
	}

	response.Success(c, http.StatusCreated, "Chart created successfully", gin.H{})
}

func (h *Handler) updateChartHanlder(c *gin.Context) {
	const op = "handler.Handler.updateChartHanlder"

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

	var chartUpdateDto dto.ChartUpdateDto
	if err = c.ShouldBindJSON(&chartUpdateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrVisualizationInvalidRequestData.Error(), err)
		return
	}

	if err = h.services.Chart.Update(c.Request.Context(), canvasID, chartUpdateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToUpdateVisualization.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Canvas updated successfully", nil)
}

func (h *Handler) deleteChartHanlder(c *gin.Context) {
	const op = "handler.Handler.deleteChartHanlder"

	chartIdStr := c.Param("id")
	if chartIdStr == "" {
		h.log.Error(fmt.Sprintf("%s: %v", op, ErrChartIdMissing))
		response.Error(c, http.StatusBadRequest, ErrChartIdMissing.Error(), nil)
		return
	}

	chartID, err := uuid.Parse(chartIdStr)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrInvalidChartId.Error(), err)
		return
	}

	if err = h.services.Chart.Delete(c.Request.Context(), chartID); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToDeleteChart.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Chart deleted successfully", nil)
}
