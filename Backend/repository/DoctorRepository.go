package repository

import (
	"api/db"
	"api/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func GetDoctorById(id int64) (doctor models.Doctor, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	conn.First(&doctor, id)

	if doctor.IdDoctor == 0 {
		doctor = models.Doctor{}
	}

	return
}

func PostDoctor(doctor models.Doctor) (err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	conn.Create(doctor)

	return
}

func VerifyDoctorByCrm(crm string) (found bool) {
	conn, err := db.OpenConnection()
	if err != nil {
		return false
	}

	var doctor models.Doctor
	if err := conn.Where("crm = ?", crm).First(&doctor).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}

	return true
}

func GetAllDoctor() (doctors []models.Doctor, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	conn.Find(&doctors)

	return
}

func PutDoctor(doctor models.Doctor) (err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	result := conn.Where("idmedico = ?", doctor.IdDoctor).Updates(models.Doctor{Name: doctor.Name, Crm: doctor.Crm, IdCbo: doctor.IdCbo})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("Nenhum dado foi atualizado")
	}

	return nil
}
