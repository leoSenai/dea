package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("123")

type Claims struct {
	User string
	jwt.StandardClaims
}

func GenerateTokenJWT(User string) (string, error) {
	claims := Claims{
		User: User,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(76 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			ReturnResponseJSON(w, http.StatusUnauthorized, "Token não Informado!", "Token not informed")
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			ReturnResponseJSON(w, http.StatusUnauthorized, "Token Inválido!", err.Error())
			return
		}

		if !token.Valid {
			ReturnResponseJSON(w, http.StatusUnauthorized, "Token Inválido!", "Token not valid")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func RefreshToken(oldTokenString string) (string, error) {
	oldToken, err := jwt.ParseWithClaims(oldTokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", err
	}

	// Extract the claims from the old token
	claims, ok := oldToken.Claims.(*Claims)
	if !ok {
		return "", fmt.Errorf("Falha ao extrair inforações do Token!")
	}

	newTokenString, err := GenerateTokenJWT(claims.User)
	if err != nil {
		return "", err
	}

	return newTokenString, nil
}
