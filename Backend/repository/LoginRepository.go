package repository

import (
	"api/models"
	"api/utils"
	"errors"
)

func VerifyLoginCredentials(login models.Login) (err error) {

	isValid := utils.IsPasswordValid(login.User, login.Password)

	if !isValid {
		err = errors.New("as credenciais sao invalidas")
	}

	return err
}