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

		templates := api.Group("/templates")
		{
			templates.POST("", h.createTemplate)
			templates.GET("", h.getAllTemplates)
			templates.GET("/:id", h.getTemplateByID)
			templates.PATCH("/:id", h.updateTemplate)
		}

	}

	return handler
}
