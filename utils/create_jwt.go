package utils

import (
	"os"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
)

func CreateJwt(userId uuid.UUID) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_TOKEN")))
	if err != nil {
		return "", err
	}
	return token, nil
}
