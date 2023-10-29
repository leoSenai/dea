package repository

import (
	"api/db"
	"api/models"
	"errors"

	"gorm.io/gorm"
)

func PostProximity(proximity models.Proximity) (err error) {
	conn, err := db.GetDB()

	if err != nil {
		return
	}

	conn.Create(proximity)

	return
}

func GetProximityByIdPersonAndIdPatient(proximity models.Proximity) (found bool) {
	conn, err := db.GetDB()

	if err != nil {
		return
	}

	if err := conn.Where("paciente_idpaciente = ? AND pessoa_idpessoa = ?", proximity.IdPatient, proximity.IdPerson).First(&proximity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}

	return true
}

func GetProximityAllByIdPerson(id int64) ([]models.Proximity, error) {
	conn, err := db.GetDB()
	if err != nil {
		return nil, err
	}

	var proximitys []models.Proximity
	if err := conn.Where("pessoa_idpessoa = ?", id).Find(&proximitys).Error; err != nil {
		return nil, err
	}

	return proximitys, nil
}

func GetProximityAllByIdPatient(id int64) ([]models.Proximity, error) {
	conn, err := db.GetDB()
	if err != nil {
		return nil, err
	}

	var proximitys []models.Proximity
	if err := conn.Where("paciente_idpaciente = ?", id).Find(&proximitys).Error; err != nil {
		return nil, err
	}

	return proximitys, nil
}
