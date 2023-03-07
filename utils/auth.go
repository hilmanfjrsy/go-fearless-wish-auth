package utils

import (
	"errors"
	"go-todo-app/model"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type JWTClaims struct {
	ID        uint  `json:"sub,omitempty"`
	ExpiresAt int64 `json:"exp,omitempty"`
	IssuedAt  int64 `json:"iat,omitempty"`
}

func (c JWTClaims) Valid(helper *jwt.ValidationHelper) (err error) {
	if helper.After(time.Unix(c.ExpiresAt, 0)) {
		err = errors.New("token has expired")
	}
	if helper.Before(time.Unix(c.IssuedAt, 0)) {
		err = errors.New("token used before issued")
	}
	return err
}

func GenerateToken(user model.User) (string, error) {
	now := time.Now()
	expiry := time.Now().Add(time.Hour * 24 * 2)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		ID:        user.ID,
		ExpiresAt: expiry.Unix(),
		IssuedAt:  now.Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
