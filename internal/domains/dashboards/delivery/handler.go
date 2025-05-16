package delivery

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"visualizer-go/internal/domains/dashboards"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type dashboardHandler struct {
	log         *slog.Logger
	dashboardUC dashboards.Usecase
}

func NewDashboardHandler(log *slog.Logger, dashboardUC dashboards.Usecase) dashboards.Handler {
	return &dashboardHandler{log: log, dashboardUC: dashboardUC}
}

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

// createDashboard godoc
//
// @Summary Create dashboard
// @Tags Dashboards
// @Router /dashboards [post]
func (h *dashboardHandler) HandleCreate(c *gin.Context) {
	const op = "handler.Handler.handleCreateDashboard"

	var visualizationCreateDto dto.VisualizationCreateDto
	if err := c.ShouldBindJSON(&visualizationCreateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, ErrVisualizationInvalidRequestData.Error(), err)
		return
	}

	templateID, err := h.dashboardUC.Create(c.Request.Context(), visualizationCreateDto)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToCreateVisualization.Error(), err)
		return
	}

	response.Success(c, http.StatusCreated, "Visualization created successfully", templateID)
}

// getDashboards godoc
//
// @Summary Get all dashboards
// @Tags Dashboards
// @Router /dashboards [get]
func (h *dashboardHandler) HandleGet(c *gin.Context) {
	const op = "handler.Handler.handleGetDashboards"

	templates, err := h.dashboardUC.GetAll(c.Request.Context())
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToFetchVisualizations.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Visualizations fetched successfully", templates)
}

// getDashboardById godoc
//
// @Summary Get dashboard by id
// @Tags Dashboards
// @Router /dashboards/{id} [get]
func (h *dashboardHandler) HandleGetById(c *gin.Context) {
	const op = "handler.Handler.handleGetDashboardByID"

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

	template, err := h.dashboardUC.GetByID(c.Request.Context(), templateID)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusNotFound, ErrVisualizationNotFound.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Visualization fetched successfully", template)
}

// updateDashboard godoc
//
// @Summary Update dashboard
// @Tags Dashboards
// @Router /dashboards [patch]
func (h *dashboardHandler) HandleUpdate(c *gin.Context) {
	const op = "handler.Handler.handleUpdateDashboard"

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

	if err = h.dashboardUC.Update(c.Request.Context(), templateID, visualizationUpdateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToUpdateVisualization.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Visualization updated successfully", nil)
}

// deleteDashboard godoc
//
// @Summary Delete dashboard
// @Tags Dashboards
// @Router /dashboards/{id} [delete]
func (h *dashboardHandler) HandleDelete(c *gin.Context) {
	const op = "handler.Handler.handleDeleteDashboard"

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

	if err = h.dashboardUC.Delete(c.Request.Context(), templateID); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToDeleteVisualization.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Visualization deleted successfully", nil)
}

// getDashboardsByTemplateId godoc
//
// @Summary Get dashboards by template id
// @Tags Dashboards
// @Router /dashboards/t/{id} [get]
func (h *dashboardHandler) HandleGetByTemplateId(c *gin.Context) {
	const op = "handler.Handler.handleGetDashboardsByTemplateID"

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

	templates, err := h.dashboardUC.GetByTemplateID(c.Request.Context(), templateID)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToFetchVisualizations.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Visualizations fetched successfully", templates)
}

// getDashboardByShareId godoc
//
// @Summary Get dashboard by share id
// @Tags Dashboards
// @Router /dashboards/share/{id} [get]
func (h *dashboardHandler) HandleGetByShareId(c *gin.Context) {
	const op = "handler.Handler.handleGetDashboardByShareID"

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

	template, err := h.dashboardUC.GetByShareID(c.Request.Context(), templateID)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusNotFound, ErrVisualizationNotFound.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Visualization fetched successfully", template)
}

// metrichDashboard godoc
//
// @Summary Metric dashboard
// @Tags Dashboards
// @Router /dashboards/{id}/metric [patch]
func (h *dashboardHandler) HandleMetric(c *gin.Context) {
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

	if err = h.dashboardUC.IncrementViewCount(c.Request.Context(), templateID); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, ErrFailedToIncrementViewCountVisualization.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "view count incremented", nil)
}
