package delivery

import (
	"visualizer-go/internal/domains/charts"

	"github.com/gin-gonic/gin"
)

func RegisterChartRoutes(g *gin.RouterGroup, h charts.Handler) {
	g.GET("", h.HandleGetByCanvasId)
	g.POST("", h.HandleCreate)
	g.PATCH("/:id", h.HandleUpdate)
	g.DELETE("/:id", h.HandleDelete)
}
