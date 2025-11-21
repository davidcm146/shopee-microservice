package main

import (
	"github.com/davidcm146/shopee-microservice/product-service/internal/server"
	"github.com/davidcm146/shopee-microservice/product-service/pkg/logger"
	"log"
)

func main() {
	logr, err := logger.NewLogger("PRODUCT_SERVICE", "development")
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}
	defer logr.Sync()
	app.Start(logr)
}
