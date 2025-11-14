package middleware

import (
	"fmt"
	"net/http"

	"github.com/davidcm146/shopee-microservice/user-service/internal/dto"
	"github.com/davidcm146/shopee-microservice/user-service/internal/service"
	"github.com/gin-gonic/gin"
)

func AuthRequired(authService service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("session_id")
		fmt.Println("Session: ", cookie.Value)
		if err != nil || cookie.Value == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		// fmt.Println("Raw Cookie Header:", c.Request.Header.Get("Cookie"))
		user, err := authService.GetUserBySessionID(c.Request.Context(), cookie.Value)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
			return
		}

		authUser := &dto.AuthResponse{
			ID:    user.ID,
			Role:  user.Role,
			Email: user.Email,
		}

		c.Set("currentUser", authUser)
		c.Next()
	}
}
