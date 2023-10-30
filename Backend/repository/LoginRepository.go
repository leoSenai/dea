package repository

import (
	"api/models"
	"api/utils"
	"errors"
)

func VerifyLoginCredentials(login models.Login) (token string, err error) {

	token, isValid := utils.IsPasswordValid(login.User, login.Password)

	if !isValid {
		err = errors.New("as credenciais sao invalidas")
	}

	return token, err
}
