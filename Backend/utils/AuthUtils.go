package utils

import (
	"api/configs"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("123")

var acessRoles = configs.GetRoles()

type Claims struct {
	User     string
	Username string
	Typeuser string
	UserId   int
	jwt.StandardClaims
}

func GenerateTokenJWT(User string, Username string, Typeuser string, UserId int) (string, error) {
	claims := Claims{
		User:     User,
		Username: Username,
		Typeuser: Typeuser,
		UserId:   UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
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
		} else {
			claims, _ := token.Claims.(*Claims)
			urlAddress := r.URL.String()
			urlAddressCut := urlAddress[1:]

			canAcess := verifyPermissionToURL(urlAddressCut, claims.Typeuser, acessRoles)

			if !canAcess {
				ReturnResponseJSON(w, http.StatusUnauthorized, "Permissão Inválida!", "Token not valid")
				return
			}

		}

		// If the token is valid, proceed to the next handler.
		next.ServeHTTP(w, r)
	})
}

func verifyPermissionToURL(path string, typeUser string, acessRoles [][]string) (canAcess bool) {
	//Esse método verifica se o usuário(termo genérico) logado tem permissão para fazer a request.

	for _, role := range acessRoles {
		if role[0] == typeUser {
			roleURLs := strings.Split(role[1], ";")
			for _, roleURL := range roleURLs {
				if strings.Contains(path, roleURL) {
					canAcess = true
					return canAcess
				}
			}
		}
	}

	canAcess = false
	return canAcess
}
