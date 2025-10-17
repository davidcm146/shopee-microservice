package middleware

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func RoleRequired(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser, exists := c.Get("currentUser")

		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		user, ok := currentUser.(struct {
			ID   string `json:"id"`
			Role string `json:"role"`
		})

		fmt.Println("Current User Role:", user)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type"})
			return
		}
		if slices.Contains(roles, user.Role) {
			c.Next()
			return
		}
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		c.Abort()
	}
}
