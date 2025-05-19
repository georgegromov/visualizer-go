package delivery

import (
	"fmt"
	"log/slog"
	"net/http"
	"visualizer-go/internal/domains/measurements"
	"visualizer-go/internal/response"
	"visualizer-go/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type measurementHandler struct {
	log           *slog.Logger
	measurementUC measurements.Usecase
}

func NewCanvasHandler(log *slog.Logger, measurementUC measurements.Usecase) measurements.Handler {
	return &measurementHandler{log: log, measurementUC: measurementUC}
}

func (h *measurementHandler) HandleGetByChartID(c *gin.Context) {

	const op = "handler.measurementHandler.HandleGetByChartID"

	cid, err := uuid.Parse(c.Query("chartId"))
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	canvases, err := h.measurementUC.GetByChartID(c.Request.Context(), cid)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Charts fetched successfully", canvases)

}
func (h *measurementHandler) HandleCreate(c *gin.Context) {
	const op = "handler.measurementHandler.HandleCreate"

	input := &measurements.MeasurementCreateDTO{}

	if err := utils.ReadRequestBody(c, input); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	err := h.measurementUC.Create(c.Request.Context(), input)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	response.Success(c, http.StatusCreated, "Canvas created successfully", gin.H{})
}
func (h *measurementHandler) HandleUpdate(c *gin.Context) {
	const op = "handler.measurementHandler.HandleUpdate"

	mid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	input := &measurements.MeasurementUpdateDTO{}

	if err := utils.ReadRequestBody(c, &input); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	err = h.measurementUC.Update(c.Request.Context(), mid, input)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Canvas updated successfully", gin.H{})
}
func (h *measurementHandler) HandleDelete(c *gin.Context) {
	const op = "handler.measurementHandler.HandleDelete"

	mid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	if err = h.measurementUC.Delete(c.Request.Context(), mid); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Canvas deleted successfully", nil)
}
