package routes

import (
	"github.com/davidcm146/shopee-microservice/user-service/internal/common/errors"
	"github.com/davidcm146/shopee-microservice/user-service/internal/handler"
	"github.com/davidcm146/shopee-microservice/user-service/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(authHandler *handler.AuthHandler) *gin.Engine {
	router := gin.Default()
	router.Use(Errors.ErrorHandler())
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.GET("/verify", authHandler.VerifySession)
			auth.Use(middleware.AuthRequired(authHandler.AuthService()))
			private := auth.Group("/")
			private.Use(middleware.AuthRequired(authHandler.AuthService()))
			{
				auth.GET("/me", authHandler.Me)
				auth.DELETE("/logout", authHandler.Logout)
			}
		}
	}
	return router
}
