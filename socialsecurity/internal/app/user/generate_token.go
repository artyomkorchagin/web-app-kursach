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

func (s *Service) GenerateToken(ctx context.Context, email, password string) (string, error) {

	user, err := s.repo.GetUser(ctx, email)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.UserID,
	})

	return token.SignedString([]byte(signingKey))
}
