package service

import (
	"context"

	"github.com/Artemych91/recipe-app/internal/models"
	"github.com/Artemych91/recipe-app/internal/repository"
)

const PER_PAGE = 10

type RecipeService struct {
	recipeRepository repository.RecipeRepository
}

func NewRecipeService(recipeRepository repository.RecipeRepository) *RecipeService {
	return &RecipeService{recipeRepository: recipeRepository}
}

func (s *RecipeService) CreateRecipe(ctx context.Context, recipe models.Recipe) (int, error) {
	// Implement the logic to create a recipe.
	// You can use the RecipeRepository to interact with the database.

	recipeId, err := s.recipeRepository.CreateRecipe(ctx, recipe)
	if err != nil {
		return 0, err
	}

	return recipeId, nil
}

func (s *RecipeService) GetRecipeByID(ctx context.Context, id int) (models.Recipe, error) {
	// Implement the logic to retrieve a recipe by ID.
	// You can use the RecipeRepository to interact with the database.

	recipe, err := s.recipeRepository.GetRecipeByID(ctx, id)
	if err != nil {
		return models.Recipe{}, err
	}

	return recipe, nil
}

func (s *RecipeService) GetRecipesByUserID(ctx context.Context, userID int) ([]models.Recipe, error) {

	recipes, err := s.recipeRepository.GetRecipesByUserID(ctx, userID)
	if err != nil {
		return make([]models.Recipe, 0), err
	}

	return recipes, nil
}

func (s *RecipeService) GetRecipesList(ctx context.Context, lastShowedRecipeTS string) ([]models.Recipe, error) {

	recipes, err := s.recipeRepository.GetRecipesList(ctx, lastShowedRecipeTS, PER_PAGE)
	if err != nil {
		return make([]models.Recipe, 0), err
	}

	return recipes, nil
}
