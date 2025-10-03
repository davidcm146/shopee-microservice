package routes

import (
	"github.com/davidcm146/shopee-microservice/user-service/internal/handler"
	"github.com/davidcm146/shopee-microservice/user-service/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(authHandler *handler.AuthHandler) *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.Use(middleware.AuthRequired(authHandler.AuthService()))
		private := auth.Group("/")
		private.Use(middleware.AuthRequired(authHandler.AuthService()))
		{
			auth.DELETE("/logout", authHandler.Logout)
			auth.GET("/me", authHandler.Me)
		}

	}
	return router
}
