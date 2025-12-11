package jwtadapter

import (
	"os"
	"time"

	"github.com/FerrySDN/auth-service/internal/ports"
	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	secret []byte
}

func NewJWTService() ports.TokenService {
	sec := os.Getenv("JWT_SECRET")
	if sec == "" {
		sec = "CHANGE_ME_DEV_ONLY"
	}
	return &jwtService{secret: []byte(sec)}
}

func (j *jwtService) Generate(username string,UserId int64) (string, error) {
	claims := jwt.MapClaims{
		"sub": username,
		"user_id":UserId,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(j.secret)
}

func (j *jwtService) Validate(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	sub, _ := claims["sub"].(string)
	return sub, nil
}
