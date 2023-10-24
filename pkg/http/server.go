package http

import (
	"database/sql"
	"net/http"

	"github.com/Artemych91/recipe-app/internal/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CreateServer(userService *service.UserService, recipeService *service.RecipeService, db *sql.DB) *http.Server {
	r := gin.Default()

	// Routers
	r.GET("/users/:userId/recipes", GetRecipesByUserHandler(recipeService))
	r.GET("/recipes/:recipeId", GetRecipeByIDHandler(recipeService))
	r.GET("/users/:userId", GetUserByIDHandler(userService))
	// r.POST("/albums", handlers.PostAlbum)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return server
}
