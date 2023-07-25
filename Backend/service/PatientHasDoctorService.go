package service

import (
	"api/models"
	"api/repository"
)

func GetPatientHasDoctorByPatientId(id int64) (patientHasDoctor []models.PatientHasDoctor, err error) {
	patientHasDoctor, err = repository.GetPatientHasDoctorByPatientId(int64(id))
	return patientHasDoctor, err
}

func GetPatientHasDoctorByDoctorId(id int64) (patientHasDoctor []models.PatientHasDoctor, err error) {
	patientHasDoctor, err = repository.GetPatientHasDoctorByDoctorId(int64(id))
	return patientHasDoctor, err
}

func GetAllPatientHasDoctor() (patientHasDoctors []models.PatientHasDoctor, err error) {
	patientHasDoctors, err = repository.GetAllPatientHasDoctor()
	return patientHasDoctors, err
}

func PostPatientHasDoctor(patientHasDoctorPost models.PatientHasDoctor) (patientHasDoctorBack models.PatientHasDoctor, err error) {
	patientHasDoctorBack, err = repository.PostPatientHasDoctor(patientHasDoctorPost)
	return patientHasDoctorBack, err
}
