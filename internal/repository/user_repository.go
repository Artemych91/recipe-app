package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Artemych91/recipe-app/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	GetUserByID(ctx context.Context, id int) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}

type UserRepositoryDB struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryDB{db: db}
}

func (r *UserRepositoryDB) CreateUser(ctx context.Context, user models.User) (int, error) {
	query := "INSERT INTO users (email, firstName, lastName, password) VALUES (?, ?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query, user.Email, user.FirstName, user.LastName, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *UserRepositoryDB) GetUserByID(ctx context.Context, id int) (models.User, error) {
	var user models.User
	query := "SELECT id, email, firstName, lastName, password FROM users WHERE id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}

func (r *UserRepositoryDB) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	query := "SELECT id, email, firstName, lastName, password FROM users WHERE email = ?"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}
