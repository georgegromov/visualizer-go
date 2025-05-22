package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"visualizer-go/internal/api/middlewares"
	"visualizer-go/internal/config"
	canvash "visualizer-go/internal/domains/canvases/delivery"
	canvaspg "visualizer-go/internal/domains/canvases/repository"
	canvasuc "visualizer-go/internal/domains/canvases/usecase"
	charth "visualizer-go/internal/domains/charts/delivery"
	chartpg "visualizer-go/internal/domains/charts/repository"
	chartuc "visualizer-go/internal/domains/charts/usecase"
	dashboardh "visualizer-go/internal/domains/dashboards/delivery"
	dashboardpg "visualizer-go/internal/domains/dashboards/repository"
	dashboarduc "visualizer-go/internal/domains/dashboards/usecase"
	measurementh "visualizer-go/internal/domains/measurements/delivery"
	measurementpg "visualizer-go/internal/domains/measurements/repository"
	measurementuc "visualizer-go/internal/domains/measurements/usecase"
	templateh "visualizer-go/internal/domains/templates/delivery"
	templatepg "visualizer-go/internal/domains/templates/repository"
	templateuc "visualizer-go/internal/domains/templates/usecase"
	userh "visualizer-go/internal/domains/users/delivery"
	userpg "visualizer-go/internal/domains/users/repository"
	useruc "visualizer-go/internal/domains/users/usecase"
	jwt_manager "visualizer-go/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	log        *slog.Logger
	config     *config.Config
	pgdb       *sqlx.DB
	handler    *gin.Engine
	httpServer *http.Server
}

func New(log *slog.Logger, config *config.Config, pgdb *sqlx.DB) *Server {
	return &Server{
		log:     log,
		config:  config,
		pgdb:    pgdb,
		handler: gin.New(),
	}
}

func (s *Server) Register() {

	// init repositories
	userRepository := userpg.NewUserRepo(s.log, s.pgdb)
	templateRepository := templatepg.NewTemplateRepo(s.log, s.pgdb)
	canvasRepository := canvaspg.NewCanvasRepo(s.log, s.pgdb)
	chartRepository := chartpg.NewChartRepo(s.log, s.pgdb)
	measurementRepository := measurementpg.NewMeasurementRepo(s.log, s.pgdb)
	dashboardRepository := dashboardpg.NewDashboardRepo(s.log, s.pgdb)

	// init managers
	jwtManager := jwt_manager.NewJwtManager(s.config.Jwt)

	// init usecases
	userUsecase := useruc.NewUserService(s.log, userRepository, jwtManager)
	templateUsecase := templateuc.NewTemplateService(s.log, templateRepository)
	canvasUsecase := canvasuc.NewCanvasUsecase(s.log, canvasRepository)
	chartUsecase := chartuc.NewChartService(s.log, chartRepository)
	measurementUsecase := measurementuc.NewMeasurementUsecase(s.log, measurementRepository)
	dashboardUsecase := dashboarduc.NewVisualizationService(s.log, dashboardRepository)

	// init handlers
	userHandler := userh.NewUserHandler(s.log, userUsecase)
	templateHandler := templateh.NewTemplateHandler(s.log, templateUsecase)
	canvasHandler := canvash.NewCanvasHandler(s.log, canvasUsecase)
	chartHandler := charth.NewChartHandler(s.log, chartUsecase)
	measurementHandler := measurementh.NewCanvasHandler(s.log, measurementUsecase)
	dashboardHandler := dashboardh.NewDashboardHandler(s.log, dashboardUsecase)

	s.handler.Use(gin.Recovery(), gin.Logger(), middlewares.CorsMiddleware(s.config.Origin))

	api := s.handler.Group("/api")
	{
		// get /api/status
		api.GET("/status", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})
		// define group route /api/auth
		auth := api.Group("/auth")
		{
			auth.POST("/login", userHandler.HandleLogin)
		}
		// define group route protected
		protected := api.Group("")
		protected.Use(middlewares.AuthMiddleware(s.log, userUsecase, jwtManager))
		{
			// define user group route /api/users
			users := protected.Group("/users")
			{
				users.GET("/me", userHandler.HandleGetMe)
				users.GET("/:id", userHandler.HandleGetById)
				users.POST("", userHandler.HandleCreate)
				users.PATCH("/:id", userHandler.HandleUpdate)
			}
			// define user group route /api/templates
			templates := protected.Group("/templates")
			{
				templates.POST("", templateHandler.HandleCreate)
				templates.GET("", templateHandler.HandleGet)
				templates.GET("/:id", templateHandler.HandleGetById)
				templates.PATCH("/:id", templateHandler.HandleUpdate)
				templates.POST("/analysis", templateHandler.HandleSaveAs)
			}
			// define canvas group route /api/canvases
			canvases := protected.Group("/canvases")
			{
				canvases.GET("", canvasHandler.HandleGetByTemplateId)
				canvases.POST("", canvasHandler.HandleCreate)
				canvases.PATCH("/:id", canvasHandler.HandleUpdate)
				canvases.DELETE("/:id", canvasHandler.HandleDelete)
			}
			// define chart group route /api/charts
			charts := protected.Group("/charts")
			{
				charts.GET("", chartHandler.HandleGetByCanvasId)
				charts.POST("", chartHandler.HandleCreate)
				charts.PATCH("/:id", chartHandler.HandleUpdate)
				charts.DELETE("/:id", chartHandler.HandleDelete)
			}
			// define chart group route /api/measurements
			measurements := protected.Group("/measurements")
			{
				measurements.GET("", measurementHandler.HandleGetByChartID)
				measurements.POST("", measurementHandler.HandleCreate)
				measurements.PATCH("/:id", measurementHandler.HandleUpdate)
				measurements.DELETE("/:id", measurementHandler.HandleDelete)
			}
			// TODO: переделать в dashboards
			// define user group route /api/dashboards
			dashboards := protected.Group("/dashboards")
			{
				dashboards.POST("", dashboardHandler.HandleCreate)
				dashboards.GET("", dashboardHandler.HandleGet)
				dashboards.GET("/:id", dashboardHandler.HandleGetById)
				dashboards.PATCH("/:id", dashboardHandler.HandleUpdate)
				dashboards.DELETE("/:id", dashboardHandler.HandleDelete)
				// переделать в api/templates/{id}/dashboards
				dashboards.GET("/t/:id", dashboardHandler.HandleGetByTemplateId)
			}

			// get /api/dashboards/:id/metrics
			api.PATCH("/dashboards/:id/metrics", dashboardHandler.HandleMetrics)
			// get /api/dashboards/share/:id
			api.GET("/dashboards/share/:id", dashboardHandler.HandleGetByShareId)
		}
	}

	s.handler.GET("/swagger/*any", func(ctx *gin.Context) {
		if ctx.Request.RequestURI == "/swagger/" {
			ctx.Redirect(302, "/swagger/index.html")
		}
		ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8888/swagger/doc.json"))(ctx)
	})

	s.httpServer = &http.Server{
		Addr:           ":" + s.config.Server.Port,
		Handler:        s.handler,
		ReadTimeout:    s.config.Server.ReadTimeout,
		WriteTimeout:   s.config.Server.WriteTimeout,
		IdleTimeout:    s.config.Server.IdleTimeout,
		MaxHeaderBytes: s.config.Server.MaxHeaderMegabytes << 20,
	}
}

func (s *Server) MustStart() {
	const op = "server.Run"

	log := s.log.With(slog.String("op", op), slog.String("port", s.config.Server.Port))

	log.Info("starting http server...", slog.String("addr", s.config.Server.Host+":"+s.config.Server.Port))

	if err := s.httpServer.ListenAndServe(); err != nil {
		panic(fmt.Errorf("%s: %w", op, err))
	}

}

func (s *Server) Stop(ctx context.Context) error {
	const op = "server.Shutdown"

	log := s.log.With(slog.String("op", op))

	log.Info("stopping http server...")

	return s.httpServer.Shutdown(ctx)
}
