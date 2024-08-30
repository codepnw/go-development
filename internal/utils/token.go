package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/codepnw/godevelopment/internal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CustomClaims struct {
	User models.User `json:"user"`
	jwt.StandardClaims
}

func GenerateJWT(user models.User, secret string) (string, error) {
	claims := CustomClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "https://github.com/codepnw/go-development",
			Audience: "https://github.com/codepnw/go-development",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateJWT(tokenString, secret string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token string")
	}
	return claims, nil
}

func GetToken(c *gin.Context) (string, error) {
	token := c.Request.Header.Get("Authorization")
	access_token := strings.TrimPrefix(token, "Bearer ")

	if access_token == "" || access_token == token {
		return "", errors.New("token missing in header of request")
	}
	return access_token, nil
}