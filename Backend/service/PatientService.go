package service

import (
	"api/models"
	"api/repository"
)

func GetPatientById(id int64) (patient models.Patient, err error) {
	patient, err = repository.GetPatientById(int64(id))
	return patient, err
}

func GetPatientByName(name string) (patient models.Patient, err error) {
	patient, err = repository.GetPatientByName(name)
	return patient, err
}

func GetAllPatient() (patients []models.Patient, err error) {
	patients, err = repository.GetAllPatient()
	return patients, err
}

func PostPatient(patientPost models.Patient) (patientBack models.Patient, err error) {
	patientBack, err = repository.PostPatient(patientPost)
	return patientBack, err
}

func PutPatient(patientPut models.Patient) (patientBack models.Patient, err error) {
	patientBack, err = repository.PutPatient(patientPut)
	return patientBack, err
}
