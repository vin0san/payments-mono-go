package http

import (
	"context"
	"fmt"
	"time"

	"pye/internal/app"
	"pye/internal/config"
	"pye/internal/repository"
	"pye/pkg/logger"
	"pye/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Server struct {
	Router      *gin.Engine
	Config      *config.Config
	DB          *pgxpool.Pool
	UserHandler *UserHandler
}

func NewServer() *Server {
	cfg := config.LoadConfig()

	db, err := repository.NewPostgres(cfg.DB)
	if err != nil {
		logger.Log.Fatal("Failed to connect to database", zap.Error(err))
	}

	if cfg.Env != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(RequestID())

	// dependencies
	userRepo := repository.NewUserRepository(db)
	userService := app.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	// build server
	s := &Server{
		Router:      router,
		Config:      cfg,
		DB:          db,
		UserHandler: userHandler,
	}

	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	// public
	s.Router.GET("/health", s.healthHandler)

	s.Router.POST("/users", s.UserHandler.Create)

	s.Router.POST("/auth/register", s.UserHandler.Register)
	s.Router.POST("/auth/login", s.UserHandler.Login)

	// protected
	protected := s.Router.Group("/")
	protected.Use(AuthMiddleware())

	protected.GET("/me", s.meHandler)
}

func (s *Server) healthHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := s.DB.Ping(ctx); err != nil {
		response.Error(c, 500, "db_error", "database unreachable", "")
		return
	}

	response.Success(c, 200, gin.H{
		"status":    "ok",
		"service":   "pye",
		"version":   "0.1.0",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%s", s.Config.ServerPort)
	logger.Log.Info("Server starting", zap.String("addr", addr))
	return s.Router.Run(addr)
}

func (s *Server) meHandler(c *gin.Context) {
	userID := c.GetString("user_id")

	response.Success(c, 200, gin.H{
		"user_id": userID,
	})
}
