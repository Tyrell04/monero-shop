package service

import (
	"context"
	"monero-shop-api/internal/core/port"
	"monero-shop-api/internal/core/util"
	"monero-shop-api/internal/exception"
)

// AuthService represents the auth service
type AuthService struct {
	repo port.UserRepository
	ts   port.TokenService
}

// NewAuthService creates a new auth service instance
func NewAuthService(repo port.UserRepository, ts port.TokenService) *AuthService {
	return &AuthService{
		repo,
		ts,
	}
}

// Login gives a registered user an access token if the credentials are valid
func (as *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := as.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", exception.ErrInvalidCredentials
	}

	match, err := util.ComparePasswordAndHash(user.Password, password)
	if err != nil {
		return "", err
	}
	if match {
		return "", exception.ErrInvalidCredentials
	}

	accessToken, err := as.ts.CreateToken(user)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
