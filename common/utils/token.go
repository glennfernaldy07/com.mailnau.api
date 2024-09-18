package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var secretKey = []byte("secret-key")

func CreateToken(key string, tokenExpTime int) (string, error) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Now().In(loc).Format(time.RFC3339)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"key": key,
			"exp": time.Now().In(loc).Add(time.Duration(tokenExpTime) * time.Second).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return claims, err
	}

	if !token.Valid {
		return claims, fmt.Errorf("invalid token")
	}

	//for key, val := range claims {
	//	if key == "key" {
	//		fmt.Printf(val.(string))
	//	}
	//}
	return claims, nil
}
