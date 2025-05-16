package delivery

import (
	"visualizer-go/internal/domains/templates"

	"github.com/gin-gonic/gin"
)

func RegisterTemplateRoutes(g *gin.RouterGroup, h templates.Handler) {
	g.POST("", h.HandleCreate)
	g.GET("", h.HandleGet)
	g.GET("/:id", h.HandleGetById)
	g.PATCH("/:id", h.HandleUpdate)
}
