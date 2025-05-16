package delivery

import (
	"visualizer-go/internal/domains/dashboards"

	"github.com/gin-gonic/gin"
)

func RegisterDashboardRoutes(g *gin.RouterGroup, h dashboards.Handler) {
	g.POST("", h.HandleCreate)
	g.GET("", h.HandleGet)
	g.GET("/:id", h.HandleGetById)
	g.PATCH("/:id", h.HandleUpdate)
	g.DELETE("/:id", h.HandleDelete)
	// переделать в api/templates/{id}/dashboards
	g.GET("/t/:id", h.HandleGetByTemplateId)
	g.PATCH("/:id/metric", h.HandleMetric)
}
