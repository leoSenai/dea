package service

import (
	"api/models"
	"api/repository"
)

func GetPersonById(id int64) (pessoa models.Person, err error) {
	pessoa, err = repository.Get(int64(id))
	return pessoa, err
}
