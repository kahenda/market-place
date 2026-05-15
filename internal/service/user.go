package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kahenda/marketplace/internal/models"
	"github.com/kahenda/marketplace/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) Register(req *models.RegisterRequest) (*models.User, error) {
	existing, err := s.Repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("email already in use")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hash),
		Location:     req.Location,
	}

	if err := s.Repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(req *models.LoginRequest) (string, error) {
	// Find user by email
	user, err := s.Repo.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("invalid email or password")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
