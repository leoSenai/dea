package repository

import (
	"api/db"
	"api/models"
	"log"
)

func GetPatientHasDoctorByPatientId(id int64) (patientHasDoctor []models.PatientHasDoctor, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.First(&patientHasDoctor, "patient_idpatient = ?", id)
	log.Printf("row: %v", row)

	return
}

func GetPatientHasDoctorByDoctorId(id int64) (patientHasDoctor []models.PatientHasDoctor, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.First(&patientHasDoctor, "medico_idmedico = ?", id)
	log.Printf("row: %v", row)

	return
}

func GetAllPatientHasDoctor() (patientHasDoctors []models.PatientHasDoctor, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	rows := conn.Find(&patientHasDoctors)
	log.Printf("rows: %v", rows)

	return
}

func PostPatientHasDoctor(patientHasDoctorPost models.PatientHasDoctor) (patientHasDoctorBack models.PatientHasDoctor, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.Create(&patientHasDoctorPost)
	log.Printf("row: %v", row)
	conn.First(&patientHasDoctorBack, patientHasDoctorPost.IdPatient)

	return
}
