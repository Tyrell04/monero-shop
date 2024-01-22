package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"monero-shop-api/internal/adapter/config"
	"monero-shop-api/internal/core/domain"
	"monero-shop-api/internal/exception"
	"time"
)

type Jwt struct {
	config *config.Jwt
}

func NewJwt(config *config.Jwt) *Jwt {
	return &Jwt{
		config,
	}
}

func (j *Jwt) CreateToken(user *domain.User) string {
	jwtSecret := j.config.Secret

	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Minute * time.Duration(j.config.AccessTokenExpireDuration)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exception.PanicLogging(err)

	return tokenSigned
}

func (j *Jwt) AuthService(token string) (*domain.TokenPayload, error) {
	jwtSecret := j.config.Secret

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	tokenPayload := &domain.TokenPayload{
		ID: claims["id"].(uuid.UUID),
	}

	return tokenPayload, nil
}
