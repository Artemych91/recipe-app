package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Artemych91/recipe-app/internal/models"
)

type RecipeRepository interface {
	CreateRecipe(ctx context.Context, recipe models.Recipe) (int, error)
	GetRecipeByID(ctx context.Context, id int) (models.Recipe, error)
	GetRecipesByUserID(ctx context.Context, userID int) ([]models.Recipe, error)
	GetRecipesList(ctx context.Context, lastShowedRecipeTS string, perPage int) ([]models.Recipe, error)
}

type RecipeRepositoryDB struct {
	db *sql.DB
}

func NewRecipeRepository(db *sql.DB) *RecipeRepositoryDB {
	return &RecipeRepositoryDB{db: db}
}

func (r *RecipeRepositoryDB) CreateRecipe(ctx context.Context, recipe models.Recipe) (int, error) {
	currentTime := time.Now()
	query := "INSERT INTO recipes (userId, title, description, createdAt) VALUES (?, ?, ?, ?)"

	result, err := r.db.ExecContext(
		ctx,
		query,
		recipe.UserID,
		recipe.Title,
		recipe.Description,
		currentTime.Format("2006-01-02 15:04:05"),
	)
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
	query := "SELECT * FROM recipes WHERE id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&recipe.ID, &recipe.UserID, &recipe.Title, &recipe.Description, &recipe.CreatedAt)
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
	query := "SELECT * FROM recipes WHERE userId = ?"
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var recipe models.Recipe
		err := rows.Scan(&recipe.ID, &recipe.UserID, &recipe.Title, &recipe.Description, &recipe.CreatedAt)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (r *RecipeRepositoryDB) GetRecipesList(ctx context.Context, lastShowedRecipeTS string, perPage int) ([]models.Recipe, error) {
	var recipes []models.Recipe
	query := "SELECT * FROM recipes WHERE 1=1"
	args := []interface{}{}

	if lastShowedRecipeTS != "" {
		query += " AND createdAt < ?"
		args = append(args, lastShowedRecipeTS)
	}

	query += " ORDER BY createdAt DESC LIMIT ?"
	args = append(args, perPage)

	rows, err := r.db.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var recipe models.Recipe
		err := rows.Scan(&recipe.ID, &recipe.UserID, &recipe.Title, &recipe.Description, &recipe.CreatedAt)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}
