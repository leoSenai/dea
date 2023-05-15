package service

import (
	"api/models"
	"api/repository"
	"fmt"
)

func GetPersonById(id int64) (person models.Person, err error) {
	person, err = repository.GetPersonById(int64(id))
	return person, err
}

func PostPerson(person models.Person) (err error) {
	found := repository.VerifyPersonByDocument(person.DocNumber)

	if found {
		return fmt.Errorf("Pessoa com o número de documento %s já cadastrado!", person.DocNumber)
	} else {
		err = repository.PostPerson(person)
	}

	return err
}

func GetAllPerson() (persons []models.Person, err error) {
	persons, err = repository.GetAllPerson()
	return persons, err
}

func PutPerson(personUpdate models.Person) (err error) {
	var person models.Person

	person, err = repository.GetPersonById(personUpdate.IdPerson)
	if err != nil {
		return err
	}

	if person.IdPerson == 0 {
		return fmt.Errorf("Pessoa não está cadastrada no sistema!")
	}

	err = repository.PutPerson(personUpdate)

	return
}
