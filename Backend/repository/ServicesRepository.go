package repository

import (
	"api/db"
	"api/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func GetServicesById(id int64) (services models.Services, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	conn.First(&services, id)

	if services.IdServices == 0 {
		services = models.Services{}
	}

	return
}

func PostServices(services models.Services) (err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	conn.Create(services)

	return
}

func VerifyServicesByDescription(description string) (found bool) {
	conn, err := db.OpenConnection()
	if err != nil {
		return false
	}

	var services models.Services
	if err := conn.Where("descricao = ?", description).First(&services).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}

	return true
}

func GetAllServices() (servicess []models.Services, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	conn.Find(&servicess)

	return
}

func PutServices(services models.Services) (err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	result := conn.Where("idservicos = ?", services.IdServices).Updates(models.Services{Desc: services.Desc})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("Nenhum dado foi atualizado")
	}

	return nil
}
