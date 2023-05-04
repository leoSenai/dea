package service

import (
	"api/models"
	"api/repository"
)

func GetPersonById(id int64) (person models.Person, err error) {
	person, err = repository.GetPersonById(int64(id))
	return person, err
}

func InsertPerson(person models.Person) (err error) {
	if repository.VerifyPersonByDocument(person.DocNumber) {
		return nil
	}
	err = repository.InsertPerson(person)
	return err
}
