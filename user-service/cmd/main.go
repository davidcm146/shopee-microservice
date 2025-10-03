package main

import (
	"log"

	"github.com/davidcm146/shopee-microservice/user-service/internal/server"
	"github.com/davidcm146/shopee-microservice/user-service/pkg/logger"
)

func main() {
	logr, err := logger.NewLogger("USER_SERVICE", "development")
	if err != nil {
		log.Fatalf("failed to init logger: %v", err)
	}
	defer logr.Sync()
	app.Start(logr)
}
