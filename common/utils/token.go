package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var secretKey = []byte("secret-key")

func CreateToken(key string, tokenExpTime int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"key": key,
			"exp": time.Now().Add(time.Hour * time.Duration(tokenExpTime)),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
