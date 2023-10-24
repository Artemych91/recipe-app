package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Artemych91/recipe-app/internal/models"
)

type RecipeRepository interface {
	CreateRecipe(ctx context.Context, recipe models.Recipe) (int, error)
	GetRecipeByID(ctx context.Context, id int) (models.Recipe, error)
	GetRecipesByUserID(ctx context.Context, userID int) ([]models.Recipe, error)
}

type RecipeRepositoryDB struct {
	db *sql.DB
}

func NewRecipeRepository(db *sql.DB) *RecipeRepositoryDB {
	return &RecipeRepositoryDB{db: db}
}

func (r *RecipeRepositoryDB) CreateRecipe(ctx context.Context, recipe models.Recipe) (int, error) {
	query := "INSERT INTO recipes (userId, title, description) VALUES (?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query, recipe.UserID, recipe.Title, recipe.Description)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *RecipeRepositoryDB) GetRecipeByID(ctx context.Context, id int) (models.Recipe, error) {
	var recipe models.Recipe
	query := "SELECT id, userId, title, description FROM recipes WHERE id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&recipe.ID, &recipe.UserID, &recipe.Title, &recipe.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return recipe, errors.New("recipe not found")
		}
		return recipe, err
	}

	return recipe, nil
}

func (r *RecipeRepositoryDB) GetRecipesByUserID(ctx context.Context, userID int) ([]models.Recipe, error) {
	var recipes []models.Recipe
	query := "SELECT id, userId, title, description FROM recipes WHERE userId = ?"
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var recipe models.Recipe
		err := rows.Scan(&recipe.ID, &recipe.UserID, &recipe.Title, &recipe.Description)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}
