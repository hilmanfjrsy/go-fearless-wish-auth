package middleware

import (
	"encoding/json"
	"fmt"
	"go-todo-app/utils"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
	c.Set("user_id", claims.UserId)
	c.Next()
}

func parseToken(tokenString, secret string) (claims utils.AuthClaims, err error) {
	decodedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return utils.AuthClaims{}, err
	}

	if claims, ok := decodedToken.Claims.(jwt.MapClaims); ok && decodedToken.Valid &&
		claims.VerifyAudience("fearless-wish-auth", true) &&
		claims.VerifyExpiresAt(time.Now().Unix(), true) &&
		claims.VerifyIssuedAt(time.Now().Unix(), true) {

		AuthClaims := utils.AuthClaims{}
		b, err := json.Marshal(claims)
		if err != nil {
			return utils.AuthClaims{}, err
		}
		err = json.Unmarshal(b, &AuthClaims)
		if err != nil {
			return utils.AuthClaims{}, err
		}
		return AuthClaims, nil
	}
	return utils.AuthClaims{}, err
}
