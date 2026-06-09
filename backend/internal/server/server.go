package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"pinger/internal/config"
)

// Server holds the dependencies for our HTTP server
type Server struct {
	config *config.Config
	// Aqui você pode adicionar as dependências futuras:
	// db             *gorm.DB
	// logger         *zerolog.Logger
	// algunService   *services.AlgumService
}

// New creates a new Server instance
func New(cfg *config.Config) *Server {
	return &Server{
		config: cfg,
	}
}

// SetupRoutes configura o roteador Gin seguindo o padrão do projeto base
func (s *Server) SetupRoutes() *gin.Engine {
	router := gin.New()

	// Middlewares globais
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(s.corsMiddleware())

	// Rota de Health Check
	router.GET("/health", s.healthCheck)

	// Grupo de rotas API v1
	api := router.Group("/api/v1")
	{
		// Exemplo de como estruturar os grupos no futuro:
		// users := api.Group("/users")
		// {
		// 	users.GET("/profile", s.getProfile)
		// }
		_ = api
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
