package handler

import (
	"net/http"

	"github.com/davidcm146/shopee-microservice/user-service/internal/dto"
	"github.com/davidcm146/shopee-microservice/user-service/internal/service"
	"github.com/davidcm146/shopee-microservice/user-service/internal/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authService service.AuthService
}

func (h *AuthHandler) AuthService() service.AuthService {
	return h.authService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

var validate = validator.New()

func (h *AuthHandler) Register(c *gin.Context) {
	var input dto.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if errs, err := validation.ValidateStruct(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	} else if len(errs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	user, err := h.authService.Register(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user, "message": "Registration successful"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input dto.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.authService.Login(c.Request.Context(), &input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Set session cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "session_id",
		Value:    result.SessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	c.JSON(http.StatusOK, gin.H{"session": result.SessionID})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	cookie, err := c.Request.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Session cookie not found"})
		return
	}

	if err := h.authService.Logout(c.Request.Context(), cookie.Value); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Clear cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func (h *AuthHandler) Me(c *gin.Context) {
	currentUser, err := c.Get("currentUser")

	if !err {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No user in context"})
		return
	}
	user := currentUser.(*dto.AuthResponse)
	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"email": user.Email,
	})
}

func (h *AuthHandler) VerifySession(c *gin.Context) {
	sessionID := c.Query("session_id")
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session_id query parameter is required"})
		return
	}

	user, err := h.authService.GetUserBySessionID(c.Request.Context(), sessionID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
	}
	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"role":  user.Role,
		"email": user.Email,
	})
}
