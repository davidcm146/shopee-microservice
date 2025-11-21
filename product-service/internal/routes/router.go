package routes

import (
	"github.com/davidcm146/shopee-microservice/product-service/internal/handler"
	"github.com/davidcm146/shopee-microservice/product-service/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(productHandler *handler.ProductHandler) *gin.Engine {
	router := gin.Default()
	router.Use(errors.ErrorHandler())

	api := router.Group("/api")
	{
		products := api.Group("/products")
		{
			products.GET("/", productHandler.GetAllProducts)
			products.GET("/:id", productHandler.GetProductByID)
			products.POST("/", middleware.AuthRequired(), middleware.RoleRequired("SELLER", "ADMIN"), productHandler.CreateProduct)
			products.PUT("/:id", middleware.AuthRequired(), middleware.RoleRequired("ADMIN"), middleware.OwnerRequired(productHandler.ProductService()), productHandler.UpdateProduct)
			products.DELETE("/:id", middleware.AuthRequired(), middleware.RoleRequired("ADMIN"), middleware.OwnerRequired(productHandler.ProductService()), productHandler.DeleteProduct)
		}
	}
	return router
}
