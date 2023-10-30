package http

import (
	"net/http"
	"strconv"

	"github.com/Artemych91/recipe-app/internal/models"
	"github.com/Artemych91/recipe-app/internal/service"
	"github.com/gin-gonic/gin"
)

func GetRecipesByUserHandler(recipeService *service.RecipeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the user ID from the request URL.
		userID, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		// Call the service to get the recipes by user ID.
		recipes, err := recipeService.GetRecipesByUserID(c, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in recipes"})
			return
		}

		if len(recipes) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Recipes not found"})
			return
		}

		// Return the recipes as JSON in the response.
		c.JSON(http.StatusOK, recipes)
	}
}

func GetRecipeByIDHandler(recipeService *service.RecipeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the recipe ID from the request URL.
		recipeID, err := strconv.Atoi(c.Param("recipeId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipe ID"})
			return
		}

		// Call the service to get the recipe by recipe ID.
		recipe, err := recipeService.GetRecipeByID(c, recipeID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
			return
		}

		// Return the recipe as JSON in the response.
		c.JSON(http.StatusOK, recipe)
	}
}

func GetRecipesList(recipeService *service.RecipeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		lastShowedTS := c.DefaultQuery("lastShowedTS", "")
		recipes, err := recipeService.GetRecipesList(c, lastShowedTS)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in recipes"})
			return
		}

		if len(recipes) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Recipes not found"})
			return
		}

		// Return the recipes as JSON in the response.
		c.JSON(http.StatusOK, recipes)
	}
}

func CreateRecipe(recipeService *service.RecipeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newRecipe models.Recipe
		err := c.BindJSON(&newRecipe)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown parameters"})
		}

		recipeId, err := recipeService.CreateRecipe(c, newRecipe)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		}

		c.JSON(http.StatusOK, gin.H{"id": recipeId})
	}
}
