package util

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	User string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("secret_key")

func GenerateJWT(first_name, last_name string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		User: first_name + " " + last_name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "admin-auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (bool, error) {
	if tokenString == "" {
		return false, errors.New("token kosong")
	}

	// Parse token menggunakan secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metode signing tidak valid")
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return false, errors.New("token tidak valid")
	}

	return true, nil
}