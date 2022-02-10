package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(username string, email string, jwtKey string) (string, error) {
	expirationTime := time.Now().Add(1440 * time.Minute)
	claims := &Claims{
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
