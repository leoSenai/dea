package repository

import (
	"api/db"
	"api/models"
	"log"
)

func GetPatientHasUserByUserId(id int64) (patientsHasUser []models.PatientHasUser, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.Where("usuario_idusuario = ?", id).Find(&patientsHasUser)
	log.Printf("row: %v", row)

	return
}

func PostPatientHasUser(patientHasUser models.PatientHasUser) (err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	conn.Create(&patientHasUser)

	return
}
