package main

import (
	"log"
	"pye/internal/delivery/http"
	"pye/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	if err := logger.Init(); err != nil {
		log.Fatal("Failed to initialize logger:", err)
	}
	defer logger.Sync()

	server := http.NewServer()
	logger.Log.Info("Starting Application", zap.String("port", server.Config.ServerPort))

	if err := server.Run(); err != nil {
		logger.Log.Fatal("Server failed", zap.Error(err))
	}
}
