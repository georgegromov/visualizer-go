package delivery

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"visualizer-go/internal/domains/charts"
	"visualizer-go/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type chartHandler struct {
	log     *slog.Logger
	chartUC charts.Usecase
}

func NewChartHandler(log *slog.Logger, chartUC charts.Usecase) charts.Handler {
	return &chartHandler{log: log, chartUC: chartUC}
}

var (
	ErrCanvasIdMissing     = errors.New("canvas id is missing")
	ErrChartIdMissing      = errors.New("chart id is missing")
	ErrInvalidChartId      = errors.New("invalid chart id")
	ErrFailedToFetchCharts = errors.New("failed to fetch charts")
	ErrFailedToCreateChart = errors.New("failed to create chart")
	ErrFailedToDeleteChart = errors.New("failed to delete chart")
	ErrBadRequest          = errors.New("bad request")
)

func (h *chartHandler) HandleGetByCanvasId(c *gin.Context) {
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
		response.Error(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	charts, err := h.chartUC.GetByCanvasID(c.Request.Context(), canvasId)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToFetchCharts.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Charts fetched successfully", charts)
}

func (h *chartHandler) HandleGetByID(c *gin.Context) {
	const op = "handler.Handler.getChartByIdHandler"

	chartID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrInvalidChartId.Error(), err)
		return
	}

	chart, err := h.chartUC.GetByID(c.Request.Context(), chartID)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToFetchCharts.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Chart fetched successfully", chart)
}

func (h *chartHandler) HandleCreate(c *gin.Context) {
	const op = "handler.Handler.createChartHanlder"

	var chartCreateDto charts.ChartCreateDto
	if err := c.ShouldBindJSON(&chartCreateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrBadRequest.Error(), err)
		return
	}

	err := h.chartUC.Create(c.Request.Context(), chartCreateDto)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToCreateChart.Error(), err)
		return
	}

	response.Success(c, http.StatusCreated, "Chart created successfully", gin.H{})
}

func (h *chartHandler) HandleUpdate(c *gin.Context) {
	const op = "handler.Handler.updateChartHanlder"

	canvasIDStr := c.Param("id")
	if canvasIDStr == "" {
		h.log.Error(fmt.Sprintf("%s: %v", op, ""))
		response.Error(c, http.StatusBadRequest, "", nil)
		return
	}

	canvasID, err := uuid.Parse(canvasIDStr)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	var chartUpdateDto charts.ChartUpdateDto
	if err = c.ShouldBindJSON(&chartUpdateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	if err = h.chartUC.Update(c.Request.Context(), canvasID, chartUpdateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Canvas updated successfully", nil)
}

func (h *chartHandler) HandleDelete(c *gin.Context) {
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

	if err = h.chartUC.Delete(c.Request.Context(), chartID); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToDeleteChart.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Chart deleted successfully", nil)
}
