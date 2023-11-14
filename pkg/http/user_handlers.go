package http

import (
	"net/http"
	"strconv"

	"github.com/Artemych91/recipe-app/internal/service"
	"github.com/Artemych91/recipe-app/pkg/auth"
	"github.com/gin-gonic/gin"
)

// GetUserByIDHandler retrieves a specific user by user ID.
func GetUserByIDHandler(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		user, err := userService.GetUserByID(c, userID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func LoginHandler(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a struct to receive user login credentials from the client
		var loginRequest struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		// Bind the request body to the loginRequest struct
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Authenticate the user by verifying the email and password
		user, err := userService.GetUserByEmail(c, loginRequest.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Verify the password (you may need to use a password hashing library)
		if !verifyPassword(loginRequest.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// Generate a JWT token for the authenticated user
		token, err := auth.GenerateToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
			return
		}

		// Send the token as a response
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func verifyPassword(plainPassword, hashedPassword string) bool {
	// Implement password verification logic (e.g., bcrypt)
	return plainPassword == hashedPassword
}
