package service

import (
	"monero-shop-api/internal/core/port"
)

type UserService struct {
	repo port.UserRepository
}

// NewUserService creates a new user service instance
func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{
		repo,
	}
}
