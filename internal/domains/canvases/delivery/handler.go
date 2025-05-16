package delivery

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"visualizer-go/internal/domains/canvases"
	"visualizer-go/internal/dto"
	"visualizer-go/internal/response"
	"visualizer-go/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type canvasHandler struct {
	log      *slog.Logger
	canvasUC canvases.Usecase
}

func NewCanvasHandler(log *slog.Logger, canvasUS canvases.Usecase) canvases.Handler {
	return &canvasHandler{log: log, canvasUC: canvasUS}
}

var (
	ErrFailedToFetchCanvases = errors.New("failed to fetch canvases")
)

func (h *canvasHandler) HandleCreate(c *gin.Context) {
	const op = "handler.Handler.createCanvasHandler"

	// TODO: дописать и исправить ошибки

	var canvasCreateDto dto.CanvasCreateDto

	if err := utils.ReadRequestBody(c, &canvasCreateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	err := h.canvasUC.Create(c.Request.Context(), canvasCreateDto)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	response.Success(c, http.StatusCreated, "Canvas created successfully", gin.H{})
}

func (h *canvasHandler) HandleGetByTemplateId(c *gin.Context) {

	const op = "handler.Handler.getCanvasByTemplateIdHandler"

	templateIDStr := c.Query("templateId")
	if templateIDStr == "" {
		h.log.Error(fmt.Sprintf("%s: %v", op, ""))
		response.Error(c, http.StatusBadRequest, "", nil)
		return
	}

	templateID, err := uuid.Parse(templateIDStr)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	canvases, err := h.canvasUC.GetCanvasesByTemplateID(c.Request.Context(), templateID)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Canvases fetched successfully", canvases)
}

func (h *canvasHandler) HandleUpdate(c *gin.Context) {
	const op = "handler.Handler.updateCanvasHandler"

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

	var canvasUpdateDto dto.CanvasUpdateDto
	if err = c.ShouldBindJSON(&canvasUpdateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, err.Error(), err)
		return
	}

	if err = h.canvasUC.Update(c.Request.Context(), canvasID, canvasUpdateDto); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Canvas updated successfully", nil)
}

func (h *canvasHandler) HandleDelete(c *gin.Context) {
	const op = "handler.Handler.deleteCanvasHandler"

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

	if err = h.canvasUC.Delete(c.Request.Context(), canvasID); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, err.Error(), err)
		return
	}

	response.Success(c, http.StatusOK, "Canvas deleted successfully", nil)
}
