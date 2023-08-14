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
