package service

import (
	"api/models"
	"api/repository"
)

func GetPatientHasUserByUserId(id int64) (patientsHasUser []models.PatientHasUser, err error) {
	patientsHasUser, err = repository.GetPatientHasUserByUserId(int64(id))
	return patientsHasUser, err
}

func PostPatientHasUser(patientHasUser models.PatientHasUser) (err error) {
	err = repository.PostPatientHasUser(patientHasUser)
	return
}
