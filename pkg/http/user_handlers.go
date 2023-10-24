package http

import (
	"net/http"
	"strconv"

	"github.com/Artemych91/recipe-app/internal/service"
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
