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
				templates.GET("/:templateId", h.getTemplateByID)
				templates.PATCH("/:id", h.updateTemplate)

			}
			// define canvas group route /api/canvases
			canvases := protected.Group("/canvases")
			{
				canvases.GET("", h.getCanvasByTemplateIdHandler)
				canvases.POST("", h.createCanvasHandler)
				canvases.PATCH("/:id", h.updateCanvasHandler)
				canvases.DELETE("/:id", h.deleteCanvasHandler)

			}
			// define chart group route /api/charts
			charts := protected.Group("/charts")
			{
				charts.GET("", h.getChartsByCanvasIdHanlder)
				charts.POST("", h.createChartHanlder)
				charts.PATCH("/:id", h.updateChartHanlder)
				charts.DELETE("/:id", h.deleteChartHanlder)
			}

			// TODO: переделать в dashboards
			// define user group route /api/dashboards
			dashboards := protected.Group("/dashboards")
			{
				dashboards.POST("", h.createVisualization)
				dashboards.GET("", h.getAllVisualizations)
				// переделать в api/templates/{id}/dashboards
				dashboards.GET("/t/:id", h.getVisualizationsByTemplateID)
				dashboards.GET("/:id", h.getVisualizationByID)
				dashboards.PATCH("/:id", h.updateVisualization)
				dashboards.PATCH("/:id/metric", h.metric)
				dashboards.DELETE("/:id", h.deleteVisualization)
			}

			// get /api/dashboards/share/:id
			api.GET("/dashboards/share/:id", h.getVisualizationByShareID)
		}
	}
	return handler
}
