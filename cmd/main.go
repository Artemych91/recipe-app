package main

import (
	"database/sql"
	"log"

	"github.com/Artemych91/recipe-app/internal/repository"
	"github.com/Artemych91/recipe-app/internal/service"
	"github.com/Artemych91/recipe-app/pkg/http"
)

func main() {
	// Initialize database connection here.
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/recipe")
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	defer db.Close()

	// Create repository instances.
	userRepo := repository.NewUserRepository(db)
	recipeRepo := repository.NewRecipeRepository(db)

	// Create service instances.
	userService := service.NewUserService(userRepo)
	recipeService := service.NewRecipeService(recipeRepo)

	// Create the HTTP server.
	server := http.CreateServer(userService, recipeService, db)

	// Start the server.
	log.Fatal(server.ListenAndServe())
}
