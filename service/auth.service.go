package service

import (
	"golang-boilerplate/dto"
	"golang-boilerplate/models"
	"golang-boilerplate/repository"
	"golang-boilerplate/utils"
)

// Service interface
type AuthService interface {
	Signup(input dto.SignupInput) (*dto.UserResponse, error)
}

// Implementation
type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Signup(input dto.SignupInput) (*dto.UserResponse, error) {
	// Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
	}

	if err := s.repo.CreateUser(&user); err != nil {
		return nil, err
	}

	response := &dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	return response, nil
}
