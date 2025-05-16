package delivery

import (
	"visualizer-go/internal/domains/users"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(g *gin.RouterGroup, h users.Handler) {
	g.POST("/login", h.HandleLogin)
	g.GET("/:id", h.HandleGetById)
	g.POST("", h.HandleCreate)
	g.PATCH("/:id", h.HandleUpdate)
}
