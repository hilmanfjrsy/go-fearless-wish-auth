package middleware

import (
	"fmt"
	"go-todo-app/datatransfers"
	"go-todo-app/utils"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	if token == "" {
		c.Next()
		return
	}
	claims, err := parseToken(token, os.Getenv("JWT_SECRET"))
	if err != nil {
		utils.ResponseError(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set("user_id", claims.ID)
	c.Next()
}

func parseToken(tokenString, secret string) (claims datatransfers.JWTClaims, err error) {
	if token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}); err != nil || !token.Valid {
		return datatransfers.JWTClaims{}, fmt.Errorf("invalid token. %s", err)
	}
	return
}
