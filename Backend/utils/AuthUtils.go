package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateTokenJWT() (tokenStr string, err error) {
	var sampleSecretKey = []byte("123")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true

	//TODO: ADD ROLES TO CLAIMS(JWT)

	tokenStr, err = token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}
	return
}
