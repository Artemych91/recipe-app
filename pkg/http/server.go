package http

import (
	"database/sql"
	"net/http"

	"github.com/Artemych91/recipe-app/internal/service"
	"github.com/Artemych91/recipe-app/pkg/auth"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CreateServer(userService *service.UserService, recipeService *service.RecipeService, db *sql.DB) *http.Server {
	r := gin.Default()

	// ====Routers====

	r.GET("/users/:userId/recipes", GetRecipesByUserHandler(recipeService))
	r.GET("/recipes/:recipeId", GetRecipeByIDHandler(recipeService))
	r.GET("/users/:userId", GetUserByIDHandler(userService))
	r.GET("/recipes", GetRecipesList(recipeService))

	//Create recipe
	r.POST("/recipes", authMiddleware(userService), CreateRecipe(recipeService))

	// Auth
	// r.POST("/register", RegisterHandler)
	r.POST("/login", LoginHandler(userService))

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return server
}

// AuthMiddleware checks for a valid JWT token in the request headers.
func authMiddleware(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Verify and decode the token
		claims, err := auth.VerifyAndDecodeToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Check if the user exists and is valid
		user, err := userService.GetUserByID(c, claims.UserID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("userID", user.ID)

		c.Next()
	}
}
