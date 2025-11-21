package middleware

import (
	"github.com/davidcm146/shopee-microservice/product-service/internal/service"
	"github.com/gin-gonic/gin"
)

func OwnerRequired(productService service.ProductService) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, exists := c.Get("currentUser")

		user := currentUser.(struct {
			ID    string `json:"id"`
			Role  string `json:"role"`
			Email string `json:"email"`
		})

		if !exists {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		if user.Role != "SELLER" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
			return
		}

		productID := c.Param("id")
		product, err := productService.FindByID(c.Request.Context(), productID)

		if err != nil {
			c.AbortWithStatusJSON(404, gin.H{"error": "Product not found"})
			return
		}

		if product.SellerID != user.ID {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
			return
		}
		c.Next()
	}
}
