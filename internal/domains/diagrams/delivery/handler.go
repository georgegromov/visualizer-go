package delivery

import (
	"fmt"
	"log/slog"
	"net/http"
	"visualizer-go/internal/domains/diagrams"
	"visualizer-go/internal/response"
	"visualizer-go/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type diagramHandler struct {
	log       *slog.Logger
	diagramUC diagrams.Usecase
}

func NewDiagramHandler(log *slog.Logger, diagramUC diagrams.Usecase) diagrams.Handler {
	return &diagramHandler{log: log, diagramUC: diagramUC}
}

// createDiagram godoc
//
// @Summary Create diagram
// @Tags Diagrams
// @Router /diagrams [post]
func (h *diagramHandler) HandleCreate(ctx *gin.Context) {
	const op = "handler.diagramHandler.HandleCreate"

	input := &diagrams.DiagramCreateDTO{}

	if err := utils.ReadRequestBody(ctx, input); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusBadRequest, "invalid request body", err)
		return
	}

	templateID, err := h.diagramUC.Create(ctx.Request.Context(), input)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusInternalServerError, "failed to create diagram", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Diagram created successfully", templateID)
}

// createDiagram godoc
//
// @Summary Get all diagrams
// @Tags Diagrams
// @Router /diagrams [get]
func (h *diagramHandler) HandleGet(c *gin.Context) {
	const op = "handler.diagramHandler.HandleGet"

	templates, err := h.diagramUC.GetAll(c.Request.Context())
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, "failed to fetch diagrams", err)
		return
	}

	response.Success(c, http.StatusOK, "Diagrams fetched successfully", templates)
}

// createDiagram godoc
//
// @Summary Get diagram by id
// @Tags Diagrams
// @Router /diagrams/{id} [get]
func (h *diagramHandler) HandleGetById(c *gin.Context) {
	const op = "handler.diagramHandler.HandleGetById"

	diagramID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, "invalid diagram id param", err)
		return
	}

	diagram, err := h.diagramUC.GetByID(c.Request.Context(), diagramID)
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusNotFound, "failed to fetch diagram", err)
		return
	}

	response.Success(c, http.StatusOK, "Diagram fetched successfully", diagram)
}

// createDiagram godoc
//
// @Summary Update diagram
// @Tags Diagrams
// @Router /diagrams/{id} [patch]
func (h *diagramHandler) HandleUpdate(ctx *gin.Context) {
	const op = "handler.diagramHandler.HandleUpdate"

	diagramID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusBadRequest, "invalid diagram id param", err)
		return
	}

	input := &diagrams.DiagramUpdateDTO{}

	if err := utils.ReadRequestBody(ctx, input); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusBadRequest, "invalid request body", err)
		return
	}

	if err = h.diagramUC.Update(ctx.Request.Context(), diagramID, input); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(ctx, http.StatusInternalServerError, "failed to update diagram", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Diagram updated successfully", nil)
}

// createDiagram godoc
//
// @Summary Delete diagram
// @Tags Diagrams
// @Router /diagrams/{id} [delete]
func (h *diagramHandler) HandleDelete(c *gin.Context) {
	const op = "handler.diagramHandler.HandleDelete"

	diagramID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusBadRequest, "invalid diagram id param", err)
		return
	}

	if err = h.diagramUC.Delete(c.Request.Context(), diagramID); err != nil {
		h.log.Error(fmt.Sprintf("%s: %v", op, err))
		response.Error(c, http.StatusInternalServerError, "failed to delete diagram", err)
		return
	}

	response.Success(c, http.StatusOK, "Diagram deleted successfully", nil)
}
