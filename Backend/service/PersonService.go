package service

import (
	"api/models"
	"api/models/dtos"
	"api/repository"
	"api/utils"
	"fmt"
)

func GetPersonById(id int64) (person models.Person, err error) {
	person, err = repository.GetPersonById(int64(id))
	return person, err
}

func GetPersonNoPasswordById(id int64) (personResultDto dtos.PersonResultDTO, err error) {
	person, err := repository.GetPersonById(int64(id))
	if err != nil {
		return personResultDto, err
	}

	personResultDto = utils.ConvertPersonToPersonResultDto(person)
	return personResultDto, nil
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

		if personDto.IdPatient != 0 && personDto.DescPerson != "" {
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

func PutPerson(personUpdate dtos.PersonDTO) (err error) {
	var person models.Person

	person, err = repository.GetPersonById(personUpdate.IdPerson)
	if err != nil {
		return err
	}

	if person.IdPerson == 0 {
		return fmt.Errorf("Pessoa não está cadastrada no sistema!")
	}

	person = utils.ConvertPersonDTOToPerson(personUpdate)

	err = repository.PutPerson(person)

	if err != nil {
		return err
	}

	if personUpdate.IdPatient != 0 && personUpdate.DescPerson != "" {
		proximity := models.Proximity{IdPatient: personUpdate.IdPatient, IdPerson: personUpdate.IdPerson, Desc: personUpdate.DescPerson}
		err = repository.PostProximity(proximity)
		if err != nil {
			return fmt.Errorf("Não foi possivel cadastrar a proximidade!")
		}
	}

	return
}

func GetPersonByDocNumber(docNumber string) (person models.Person, err error) {
	person, err = repository.GetPersonByDocNumber(docNumber)
	return person, err
}

func ResetPassword(personResetPassword dtos.PersonResetPasswordDTO) (isReseted bool, err error) {
	person, err := repository.GetPersonById(personResetPassword.IdPerson)

	if err != nil {
		return false, err
	}

	password, err := utils.GenerateRandomPassword(8)

	person.Password = password

	err = repository.PutPerson(person)

	if err != nil {
		return false, err
	}

	to := person.Email

	subject := "DEA - Nova senha de acesso"
	body := "Sua nova senha: " + password + "\n\n"

	err = utils.SendEmail(to, subject, body)
	if err != nil {
		return false, fmt.Errorf("erro ao enviar e-mail de reset de senha do paciente")
	}

	return true, nil
}
