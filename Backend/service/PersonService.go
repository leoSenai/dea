package service

import (
	"api/models"
	"api/models/dtos"
	"api/repository"
	"api/service/utils"
	"fmt"
)

func GetPersonById(id int64) (person models.Person, err error) {
	person, err = repository.GetPersonById(int64(id))
	return person, err
}

func PostPerson(personDto dtos.PersonDTO) (err error) {
	found := repository.VerifyPersonByDocument(personDto.DocNumber)

	if found {
		return fmt.Errorf("Pessoa com o número de documento %s já cadastrado!", personDto.DocNumber)
	} else {

		person, err := repository.PostPerson(utils.ConvertPersonDTOToPerson(personDto))
		if err != nil {
			return fmt.Errorf("Não foi possivel cadastrar está pessoa!")
		}

		if personDto.IdPatient != 0 {
			proximity := models.Proximity{IdPatient: personDto.IdPatient, IdPerson: person.IdPerson, Desc: personDto.DescPerson}
			err = repository.PostProximity(proximity)
			if err != nil {
				return fmt.Errorf("Não foi possivel cadastrar a proximidade!")
			}
		}
	}

	return nil
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

func GetPersonByDocNumber(docNumber string) (person models.Person, err error) {
	person, err = repository.GetPersonByDocNumber(docNumber)
	return person, err
}
