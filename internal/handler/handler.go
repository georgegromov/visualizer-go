package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"visualizer-go/internal/middlewares"
	"visualizer-go/internal/service"
)

type Handler struct {
	log      *slog.Logger
	services *service.Service
	origin   string
}

func New(log *slog.Logger, service *service.Service, origin string) *Handler {
	return &Handler{
		log:      log,
		services: service,
		origin:   origin,
	}
}

func (h *Handler) Init() *gin.Engine {
	handler := gin.New()

	handler.Use(gin.Recovery(), gin.Logger(), middlewares.CorsMiddleware(h.origin))

	api := handler.Group("/api")
	{
		api.GET("/status", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})

		auth := api.Group("/auth")
		{
			auth.POST("/login", h.login)
		}

		protected := api.Group("")
		protected.Use(middlewares.AuthMiddleware(h.log))
		{
			users := protected.Group("/users")
			{
				users.GET("/:id", h.getUserByID)
				users.POST("", h.createUser)
				users.PATCH("/:id", h.updateUser)
			}

			templates := protected.Group("/templates")
			{
				templates.POST("", h.createTemplate)
				templates.GET("", h.getAllTemplates)
				templates.GET("/:id", h.getTemplateByID)
				templates.PATCH("/:id", h.updateTemplate)
			}

			visualizations := protected.Group("/visualizations")
			{
				visualizations.POST("", h.createVisualization)
				visualizations.GET("", h.getAllVisualizations)
				visualizations.GET("/:id", h.getVisualizationByID)
				visualizations.GET("/share/:id", h.getVisualizationByShareID)
				visualizations.PATCH("/:id", h.updateVisualization)
				visualizations.DELETE("/:id", h.deleteVisualization)
			}
		}

	}
	return handler
}
