package service

import (
	"api/models"
	"api/repository"
)

func GetPersonById(id int64) (pessoa models.Person, err error) {
	pessoa, err = repository.Get(int64(id))
	return pessoa, err
}

func InsertPerson(person models.Person) (err error) {
	if repository.VerifyPersonByDocument(person.DocNumber) {
		return nil
	}
	err = repository.InsertPerson(person)
	return err
}
