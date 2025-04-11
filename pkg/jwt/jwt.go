package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTConfig struct {
	SecretKey string
}

func NewJWTConfig(secretKey string) *JWTConfig {
	return &JWTConfig{SecretKey: secretKey}
}

func (c *JWTConfig) EncodeJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(c.SecretKey))
}
