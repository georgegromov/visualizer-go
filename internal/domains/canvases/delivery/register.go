package delivery

import (
	"visualizer-go/internal/domains/canvases"

	"github.com/gin-gonic/gin"
)

func RegisterCanvasRoutes(g *gin.RouterGroup, h canvases.Handler) {
	g.GET("", h.HandleGetByTemplateId)
	g.POST("", h.HandleCreate)
	g.PATCH("/:id", h.HandleUpdate)
	g.DELETE("/:id", h.HandleDelete)
}
