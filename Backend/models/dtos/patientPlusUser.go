package dtos

import "api/models"

type PatientPlusUser struct {
	Patient models.Patient
	IdUser  int64
}