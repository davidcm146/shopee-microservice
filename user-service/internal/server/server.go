package app

import (
	"github.com/davidcm146/shopee-microservice/user-service/config"
	"github.com/davidcm146/shopee-microservice/user-service/internal/handler"
	"github.com/davidcm146/shopee-microservice/user-service/internal/repository"
	"github.com/davidcm146/shopee-microservice/user-service/internal/routes"
	"github.com/davidcm146/shopee-microservice/user-service/internal/service"
	"github.com/davidcm146/shopee-microservice/user-service/pkg/db"
	"go.uber.org/zap"
)

func Start(logger *zap.Logger) {
	// Load MongoDB and Redis
	mongoCfg := db.InitMongo()
	redisCfg := db.InitRedis()

	userRepo := repository.NewUserRepository(mongoCfg)

	authService := service.NewAuthService(userRepo, redisCfg)
	authHandler := handler.NewAuthHandler(authService)
	router := routes.NewRouter(authHandler)

	port := config.LoadEnv("APP_PORT", "8080")
	logger.Info("Starting server", zap.String("port", port))
	if err := router.Run(":" + port); err != nil {
		logger.Fatal("Service failed to start", zap.Error(err))
	}
}
