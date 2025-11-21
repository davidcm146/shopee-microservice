package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/davidcm146/shopee-microservice/user-service/config"
	"github.com/davidcm146/shopee-microservice/user-service/internal/handler"
	"github.com/davidcm146/shopee-microservice/user-service/internal/repository"
	"github.com/davidcm146/shopee-microservice/user-service/internal/routes"
	"github.com/davidcm146/shopee-microservice/user-service/internal/service"
	"github.com/davidcm146/shopee-microservice/user-service/pkg/db"
	"go.uber.org/zap"
)

func Start(logger *zap.Logger) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Init DB
	mongoCfg := db.InitMongo()
	redisCfg := db.InitRedis()

	userRepo := repository.NewUserRepository(mongoCfg)
	authService := service.NewAuthService(userRepo, redisCfg)
	authHandler := handler.NewAuthHandler(authService)
	router := routes.NewRouter(authHandler)

	port := config.LoadEnv("PORT", "8080")
	logger.Info("Starting server", zap.String("port", port))

	go func() {
		if err := router.Run(":" + port); err != nil {
			logger.Fatal("Service failed to start", zap.Error(err))
		}
	}()

	// Wait signal
	<-ctx.Done()
	logger.Info("Shutting down server...")

	// Cleanup
	if err := mongoCfg.MongoClient.Disconnect(context.Background()); err != nil {
		logger.Error("Error disconnecting MongoDB", zap.Error(err))
	}
	if err := redisCfg.Client.Close(); err != nil {
		logger.Error("Error closing Redis", zap.Error(err))
	}

	logger.Info("Server stopped gracefully")
}
