package service

import (
	"api/models"
	"api/repository"
)

func PostLogin(login models.Login) (token string, err error) {

	token, err = repository.VerifyLoginCredentials(login)

	return token, err
}
