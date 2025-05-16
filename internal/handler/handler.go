package handler

import (
	"log/slog"
	"net/http"
	"visualizer-go/internal/middlewares"
	"visualizer-go/internal/service"
	jwt_manager "visualizer-go/pkg/jwt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

func (h *Handler) Init(jwtManager *jwt_manager.JwtManager) *gin.Engine {
	g := gin.New()

	g.Use(gin.Recovery(), gin.Logger(), middlewares.CorsMiddleware(h.origin))

	// define group route /api
	api := g.Group("/api")
	{
		// get /api/status
		api.GET("/status", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})

		// define group route /api/auth
		auth := api.Group("/auth")
		{
			auth.POST("/login", h.handleLogin)
		}

		// define group route protected
		protected := api.Group("")
		protected.Use(middlewares.AuthMiddleware(h.log, h.services, jwtManager))
		{
			// define user group route /api/users
			users := protected.Group("/users")
			{
				users.GET("/:id", h.handleGetUserByID)
				users.POST("", h.handleCreateUser)
				users.PATCH("/:id", h.handleUpdateUser)
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
				dashboards.POST("", h.handleCreateDashboard)
				dashboards.GET("", h.handleGetDashboards)
				dashboards.GET("/:id", h.handleGetDashboardByID)
				dashboards.PATCH("/:id", h.handleUpdateDashboard)
				dashboards.DELETE("/:id", h.handleDeleteDashboard)
				// переделать в api/templates/{id}/dashboards
				dashboards.GET("/t/:id", h.handleGetDashboardsByTemplateID)
				dashboards.PATCH("/:id/metric", h.handleMetricDashboard)
			}

			// get /api/dashboards/share/:id
			api.GET("/dashboards/share/:id", h.handleGetDashboardByShareID)
		}
	}

	g.GET("/swagger/*any", func(ctx *gin.Context) {
		if ctx.Request.RequestURI == "/swagger/" {
			ctx.Redirect(302, "/swagger/index.html")
		}
		ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8888/swagger/doc.json"))(ctx)
	})

	return g
}
