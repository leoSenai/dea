package repository

import (
	"api/db"
	"api/models"
	"api/models/dtos"
	"api/utils"
	"log"
)

func GetPatientById(id int64) (patient models.Patient, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.First(&patient, id)
	log.Printf("row: %v", row)

	return
}

func GetAllPatientsByUserID(id int64) (patients []models.Patient, err error) {
	//conn, err := db.GetDB()
	if err != nil {
		return
	}

	//conn.Where("paciente_idpaciente").Find(&patients, id)

	return patients, err
}

func GetPatientByName(name string) (patient models.Patient, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.First(&patient, "nome = ?", name)
	log.Printf("row: %v", row)

	return
}

func GetAllPatient() (patients []models.Patient, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	rows := conn.Find(&patients)
	log.Printf("rows: %v", rows)

	return
}

func PostPatient(patientPost dtos.PatientPlusUser) (patientBack models.Patient, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	passwordEncrypted, salt := utils.GenerateEncryptedPassword(patientPost.Patient.Password)
	patientPost.Patient.Password = passwordEncrypted
	patientPost.Patient.Salt = salt

	row := conn.Create(&patientPost.Patient)
	log.Printf("row: %v", row)

	conn.First(&patientBack, patientPost.Patient.IdPatient)

	row = conn.Create(&models.PatientHasUser{
		IdPatient: patientPost.Patient.IdPatient,
		IdUser:    patientPost.IdUser,
	})
	log.Printf("row: %v", row)

	return
}

func PutPatient(patientPut models.Patient) (patientBack models.Patient, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	if patientPut.Password != "" {
		passwordEncrypted, salt := utils.GenerateEncryptedPassword(patientPut.Password)
		patientPut.Password = passwordEncrypted
		patientPut.Salt = salt
	}

	if patientPut.IdPatient != 0 {
		row := conn.Table("paciente").Where("idpaciente = ?", patientPut.IdPatient).Updates(&patientPut)
		log.Printf("row: %v", row)
	}

	return
}
