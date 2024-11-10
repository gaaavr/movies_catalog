package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"web_lab/internal/models"
)

const (
	secretKey   = "secret"
	tokenExpire = 1 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Role     string `json:"role"`
	UserID   int64  `json:"user_id"`
}

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) TokenIsValid(token string) (bool, string, int64, error) {
	parseToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return false, "", 0, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := parseToken.Claims.(*tokenClaims)
	if !ok {
		return false, "", 0, errors.New("invalid claims")
	}

	if claims.ExpiresAt.After(time.Now().Add(tokenExpire)) {
		return false, "", 0, errors.New("token is expired")
	}

	return true, claims.Role, claims.UserID, nil
}

func (s *Service) GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID:   int64(user.ID),
		Username: user.Username,
		Role:     user.Role,
	})

	return token.SignedString([]byte(secretKey))
}
