package utils

import (
	"api/models"
	"api/models/dtos"
)

func ConvertPersonDTOToPerson(personDto dtos.PersonDTO) models.Person {
	person := models.Person{
		IdPerson:  personDto.IdPerson,
		Name:      personDto.Name,
		BornDate:  personDto.BornDate,
		DocNumber: personDto.DocNumber,
		DocType:   personDto.DocType,
		Password:  personDto.Password,
		Salt:      personDto.Salt,
		Email:     personDto.Email,
	}
	return person
}

func ConvertPersonToPersonResultDto(person models.Person) dtos.PersonResultDTO {
	personDto := dtos.PersonResultDTO{
		IdPerson:  person.IdPerson,
		Name:      person.Name,
		BornDate:  person.BornDate,
		DocNumber: person.DocNumber,
		DocType:   person.DocType,
		Email:     person.Email,
		Phone:     person.Phone,
	}
	return personDto
}
