package service

import (
	"context"

	"github.com/Artemych91/recipe-app/internal/models"
	"github.com/Artemych91/recipe-app/internal/models/dto"
	"github.com/Artemych91/recipe-app/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) RegisterUser(ctx context.Context, email, firstName, lastName, password string) (int, error) {

	// Hash and salt the password before storing it in the database.
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return 0, err
	}

	user := models.User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Password:  hashedPassword,
	}

	userID, err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (dto.UserDTO, error) {

	user, err := s.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return dto.UserDTO{}, err
	}

	userDTO := dto.UserDTO{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	return userDTO, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (models.User, error) {

	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	// Generate a salt and hash the password with it.
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func checkPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
