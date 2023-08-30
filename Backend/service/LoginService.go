package service

import (
	"api/models"
	"api/repository"
)

func PostLogin(login models.Login) (err error) {

	err = repository.VerifyLoginCredentials(login)

	return err
}
