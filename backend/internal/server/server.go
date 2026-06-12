package server

import (
	"net/http"

	"pinger/internal/config"
	"pinger/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Server struct {
	config         *config.Config
	db             *gorm.DB
	logger         *zerolog.Logger
	MonitorService *services.MontiorsService
}

func New(cfg *config.Config, db *gorm.DB, logger *zerolog.Logger, monitorService *services.MontiorsService) *Server {
	return &Server{
		config:         cfg,
		db:             db,
		logger:         logger,
		MonitorService: monitorService,
	}
}

func (s *Server) SetupRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(s.corsMiddleware())

	router.GET("/health", s.healthCheck)

	api := router.Group("/api/v1")
	{
		monitors := api.Group("/monitors")
		{
			monitors.GET("", s.findAll)
			monitors.POST("/create", s.create)
			monitors.POST("/:id/ping", s.ping)
			monitors.PATCH("/:id", s.update)
			monitors.DELETE("/:id", s.delete)
		}
	}

	return router
}

func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (s *Server) corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
