package service

import (
	"api/models"
	"api/repository"
	"api/utils"
	"fmt"
)

func GetPatientById(id int64) (patient models.Patient, err error) {
	patient, err = repository.GetPatientById(int64(id))
	return patient, err
}

func GetAllPatientsByUserID(id int64) (patients []models.Patient, err error) {
	patients, err = repository.GetAllPatientsByUserID(id)
	return patients, err
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

func ResetPasswordPatient(oldPatient models.Patient) (err error) {
	var password string
	password, err = utils.GenerateRandomPassword(8)
	if err != nil {
		return fmt.Errorf("erro ao gerar senha aleatória para o paciente")
	}

	oldPatient.Password = password

	newPatient, err := repository.PutPatient(oldPatient)

	if err != nil {
		return fmt.Errorf("erro ao atualizar senha do paciente")
	}

	newPatient, err = repository.GetPatientById(oldPatient.IdPatient)

	if err != nil {
		return fmt.Errorf("erro ao buscar usuario apos atualizar senha")
	}

	to := newPatient.Email

	subject := "DEA - Nova senha de acesso"
	body := "Sua nova senha: " + password + "\n\n"

	err = utils.SendEmail(to, subject, body)
	if err != nil {
		return fmt.Errorf("erro ao enviar e-mail de reset de senha do paciente")
	}

	return err
}
