package service

import (
	"api/models"
	"api/repository"
	"fmt"
)

func GetDoctorById(id int64) (doctor models.Doctor, err error) {
	doctor, err = repository.GetDoctorById(int64(id))
	return doctor, err
}

func PostDoctor(doctor models.Doctor) (err error) {
	found := repository.VerifyDoctorByCrm(doctor.Crm)

	if found {
		return fmt.Errorf("Doutor com o CRM %s já cadastrado!", doctor.Crm)
	} else {
		err = repository.PostDoctor(doctor)
	}

	return err
}

func GetAllDoctor() (doctors []models.Doctor, err error) {
	doctors, err = repository.GetAllDoctor()
	return doctors, err
}

func PutDoctor(doctorUpdate models.Doctor) (err error) {
	var doctor models.Doctor

	doctor, err = repository.GetDoctorById(doctorUpdate.IdDoctor)
	if err != nil {
		return err
	}

	if doctor.IdDoctor == 0 {
		return fmt.Errorf("Médico não está cadastrado no sistema!")
	}

	err = repository.PutDoctor(doctorUpdate)

	return
}
