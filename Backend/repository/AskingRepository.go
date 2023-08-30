package repository

import (
	"api/db"
	"api/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func GetAskingById(id int64) (asking models.Asking, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	conn.First(&asking, id)

	if asking.IdAsking == 0 {
		asking = models.Asking{}
	}

	return
}

func PostAsking(asking models.Asking) (err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	conn.Create(asking)

	return
}

func VerifyAskingByDescription(description string) (found bool) {
	conn, err := db.OpenConnection()
	if err != nil {
		return false
	}

	var asking models.Asking
	if err := conn.Where("descricao = ?", description).First(&asking).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}

	return true
}

func GetAllAsking() (askings []models.Asking, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	conn.Find(&askings)

	return
}

func PutAsking(asking models.Asking) (err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	result := conn.Where("idpergunta = ?", asking.IdAsking).Updates(models.Asking{Desc: asking.Desc})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("nenhum dado foi atualizado")
	}

	return nil
}
