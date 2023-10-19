package service

import (
	"api/models"
	"api/models/dtos"
	"api/repository"
	"fmt"
)

func PostProximity(proximity models.Proximity) (err error) {
	person, err := repository.GetPersonById(proximity.IdPerson)
	if err != nil {
		return fmt.Errorf("Erro ao buscar pessoa: %v", err)
	}

	if person.IdPerson == 0 {
		return fmt.Errorf("Pessoa não cadastrada")
	}

	patient, err := repository.GetPatientById(proximity.IdPatient)
	if err != nil {
		return fmt.Errorf("Erro ao buscar paciente: %v", err)
	}

	if patient.IdPatient == 0 {
		return fmt.Errorf("Paciente não cadastrado")
	}

	found := repository.GetProximityByIdPersonAndIdPatient(proximity)

	if found {
		return fmt.Errorf("Proximidade de pessoa e paciente já cadastrado!")
	} else {
		err = repository.PostProximity(proximity)
	}

	return err
}

func GetProximityAllByIdPerson(id int64) (proximitys []models.Proximity, err error) {
	proximitys, err = repository.GetProximityAllByIdPerson(int64(id))
	return proximitys, err
}

func GetProximityAllByIdPatient(id int64) (proximitys []models.Proximity, err error) {
	proximitys, err = repository.GetProximityAllByIdPatient(int64(id))
	return proximitys, err
}

func GetPersonProximityAllByIdPatient(id int64) (persons []dtos.PersonResultDTO, err error) {
	proximitys, err := GetProximityAllByIdPatient(int64(id))
	if err != nil {
		return nil, err
	}

	if len(proximitys) == 0 {
		return nil, fmt.Errorf("Não há proximidades cadastradas a este paciente")
	}

	for _, proximity := range proximitys {
		person, err := GetPersonNoPasswordById(proximity.IdPerson)
		if err != nil || person.IdPerson == 0 {
			return nil, fmt.Errorf("Erro ao buscar pessoa relacionada!")
		}
		persons = append(persons, person)
	}

	return persons, nil
}
