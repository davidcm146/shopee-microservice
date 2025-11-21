package middleware

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("session_id")

		if err != nil || cookie.Value == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		resp, err := http.Get("http://localhost:8080/api/auth/verify?session_id=" + cookie.Value)
		if err != nil || resp.StatusCode != http.StatusOK {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
			return
		}
		defer resp.Body.Close()
		var user struct {
			ID    string `json:"id"`
			Role  string `json:"role"`
			Email string `json:"email"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info"})
			return
		}
		c.Set("currentUser", user)
		c.Next()
	}
}
