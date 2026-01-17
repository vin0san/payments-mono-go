// internal/delivery/http/server.go
package http

import (
	"fmt"
	"pye/internal/config"
	"pye/pkg/logger"
	"pye/pkg/response"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	Router *gin.Engine
	Config *config.Config
}

func NewServer() *Server {
	cfg := config.LoadConfig()

	// In production, disable Gin's debug logs
	if cfg.Env != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New() // Use gin.New() + custom middleware instead of Default()
	router.Use(gin.Recovery())
	router.Use(RequestID()) // Custom middleware for request IDs

	s := &Server{
		Router: router,
		Config: cfg,
	}

	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	s.Router.GET("/health", healthHandler)
}

func healthHandler(c *gin.Context) {
	response.Success(c, 200, gin.H{
		"status":     "ok",
		"service":    "pye",
		"version":    "0.1.0",
		"timestamp":  time.Now().UTC().Format(time.RFC3339),
	})
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%s", s.Config.ServerPort)
	logger.Log.Info("Server starting", zap.String("addr", addr))
	return s.Router.Run(addr)
}
