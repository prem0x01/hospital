package services

import (
	"errors"

	"context"

	"github.com/prem0x01/hospital/internal/domain"
	"github.com/prem0x01/hospital/internal/repository"
	"github.com/prem0x01/hospital/internal/utils"
)

type AuthService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Login(req *domain.LoginRequest) (*domain.AuthResponse, error) {
	ctx := context.Background()
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		return nil, errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(int(user.ID), user.Role, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *AuthService) Register(req *domain.RegisterRequest) (*domain.AuthResponse, error) {
	ctx := context.Background()
	_, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:        req.Email,
		PasswordHash: hashedPassword,
		Role:         req.Role,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	token, err := utils.GenerateJWT(int(user.ID), user.Role, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}
