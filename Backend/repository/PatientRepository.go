package repository

import (
	"api/db"
	"api/models"
	"log"
)

func GetPatientById(id int64) (patient models.Patient, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	row := conn.First(&patient, id)
	log.Printf("row: %v", row)

	return
}

func GetAllPatient() (patients []models.Patient, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	rows := conn.Find(&patients)
	log.Printf("rows: %v", rows)

	return
}

func PostPatient(patientPost models.Patient) (patientBack models.Patient, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	row := conn.Create(&patientPost)
	log.Printf("row: %v", row)

	return
}

func PutPatient(patientPut models.Patient) (patientBack models.Patient, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	patientBack = patientPut
	patientBack.IdPatient = 0

	if patientPut.IdPatient != 0 {
		row := conn.Table("paciente").Where("idpaciente = ?", patientPut.IdPatient).Updates(&patientBack)
		log.Printf("row: %v", row)
	}

	return
}
