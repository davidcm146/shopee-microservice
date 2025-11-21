package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/davidcm146/shopee-microservice/product-service/config"
	"github.com/davidcm146/shopee-microservice/product-service/internal/handler"
	"github.com/davidcm146/shopee-microservice/product-service/internal/repository"
	"github.com/davidcm146/shopee-microservice/product-service/internal/routes"
	"github.com/davidcm146/shopee-microservice/product-service/internal/service"
	"github.com/davidcm146/shopee-microservice/product-service/pkg/db"
	"go.uber.org/zap"
)

func Start(logger *zap.Logger) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	mongoCfg := db.InitMongo()

	productRepo := repository.NewProductRepository(mongoCfg)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)
	router := routes.NewRouter(productHandler)

	port := config.LoadEnv("PORT", "8081")
	logger.Info("Starting server", zap.String("port", port))

	go func() {
		if err := router.Run((":" + port)); err != nil {
			logger.Fatal("Service failed to start", zap.Error(err))
		}
	}()

	<-ctx.Done()
	logger.Info("Shutting down server...")

	if err := mongoCfg.MongoClient.Disconnect(context.Background()); err != nil {
		logger.Error("Error disconnecting MongoDB", zap.Error(err))
	}

	logger.Info("Server stopped gracefully")
}
