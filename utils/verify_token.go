package utils

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_TOKEN")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
