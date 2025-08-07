package routes

import (
	"github.com/davidcm146/shopee-microservice/user-service/internal/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(authHandler *handler.AuthHandler) *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/register")
	}
}
