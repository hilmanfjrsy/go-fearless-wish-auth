package utils

import (
	"go-todo-app/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthClaims struct {
	UserId uint `json:"user_id,omitempty"`
	jwt.RegisteredClaims
}

func GenerateToken(user model.User) (string, error) {
	now := time.Now()
	expiry := time.Now().Add(time.Hour * 24 * 2)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthClaims{
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"fearless-wish-auth"},
			ExpiresAt: jwt.NewNumericDate(expiry),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
