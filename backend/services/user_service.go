package services

import (
	"backend/models"
	"backend/pkg/hash"
	"backend/pkg/jwt"
	"backend/repositories"
	"errors"
)

type UserService struct {
	userRepo repositories.UserRepositoryInterface
}

// Interface
type UserServiceInterface interface {
	Register(user *models.User) error
	Login(email, password string) (string, error)
	GetUserById(userId int64) (*models.User, error)
}

func NewUserService(userRepo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(user *models.User) error {
	// Check if username already exists
	if _, err := s.userRepo.FindByUsername(user.Username); err == nil {
		return errors.New("username already exists")
	}

	// Check if email already exists
	if _, err := s.userRepo.FindByEmail(user.Email); err == nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return s.userRepo.Create(user)
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email")
	}

	if !hash.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := jwt.GenerateToken(uint(user.ID), "your-secret-key") // TODO: Use config for secret
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) GetUserById(userId int64) (*models.User, error) {
	return s.userRepo.FindById(userId)
}
