package service

import (
	"api/models"
	"api/repository"
)

func GetPersonById(id int64) (person models.Person, err error) {
	person, err = repository.GetPersonById(int64(id))
	return person, err
}
