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

func InsertPerson(person models.Person) (err error) {
	found, err := repository.VerifyPersonByDocument(person.DocNumber)
	if err != nil {
		return err
	}

	if found {
		return fmt.Errorf("Pessoa com o número de documento %s já cadastrado!", person.DocNumber)
	}

	err = repository.InsertPerson(person)

	return err
}

func GetAllPersons() (persons []models.Person, err error) {
	persons, err = repository.GetAllPersons()
	return persons, err
}

func UpdatePerson(personUpdate models.Person) (err error) {
	var person models.Person

	person, err = repository.GetPersonById(personUpdate.IdPerson)
	if err != nil {
		return err
	}

	if person.IdPerson == 0 {
		return fmt.Errorf("Pessoa não está cadastrada no sistema!")
	}

	err = repository.UpdatePerson(personUpdate)

	return
}
