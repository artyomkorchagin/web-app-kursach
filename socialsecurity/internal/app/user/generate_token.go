package user

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

const (
	signingKey = "qrkj#%@FNSAzpZ!@M<24FjH" // i need to change it to secret
	tokenTTL   = 12 * time.Hour
)

func (s *Service) GenerateToken(ctx context.Context, username, password string) (string, error) {
	hashedpass, err := HashPassword(password)
	if err != nil {
		return "", err
	}
	user, err := s.repo.GetUser(ctx, username, hashedpass)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}
