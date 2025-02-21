package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"visualizer-go/internal/service"
)

type Handler struct {
	log      *slog.Logger
	services *service.Service
}

func New(log *slog.Logger, service *service.Service) *Handler {
	return &Handler{
		log:      log,
		services: service,
	}
}

func (h *Handler) Init() *gin.Engine {
	handler := gin.New()

	handler.Use(gin.Recovery(), gin.Logger())

	api := handler.Group("/api")
	{
		api.GET("/status", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})

	}

	return handler
}
