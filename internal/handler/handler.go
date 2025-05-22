package handler

import (
	"log/slog"
	"net/http"
	"visualizer-go/internal/middlewares"
	"visualizer-go/internal/service"

	"github.com/gin-gonic/gin"
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

	// define group route /api
	api := handler.Group("/api")
	{
		// get /api/status
		api.GET("/status", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})

		// define group route /api/auth
		auth := api.Group("/auth")
		{
			auth.POST("/login", h.login)
		}

		// get /api/visualizations/share/:id
		api.GET("/visualizations/share/:id", h.getVisualizationByShareID)
		api.PATCH("visualizations/:id/metric", h.metric)

		// define group route protected
		protected := api.Group("")
		protected.Use(middlewares.AuthMiddleware(h.log))
		{
			// define user group route /api/users
			users := protected.Group("/users")
			{
				users.GET("/:id", h.getUserByID)
				users.POST("", h.createUser)
				users.PATCH("/:id", h.updateUser)
			}
			// define user group route /api/templates
			templates := protected.Group("/templates")
			{
				templates.POST("", h.createTemplate)
				templates.GET("", h.getAllTemplates)
				templates.GET("/:id", h.getTemplateByID)
				templates.PATCH("/:id", h.updateTemplate)
			}

			// TODO: переделать в dashboards
			// define user group route /api/visualizations
			visualizations := protected.Group("/visualizations")
			{
				visualizations.POST("", h.createVisualization)
				visualizations.GET("", h.getAllVisualizations)
				// переделать в api/templates/{id}/dashboards
				visualizations.GET("/t/:id", h.getVisualizationsByTemplateID)
				visualizations.GET("/:id", h.getVisualizationByID)
				visualizations.PATCH("/:id", h.updateVisualization)
				visualizations.DELETE("/:id", h.deleteVisualization)
			}
		}
	}
	return handler
}
