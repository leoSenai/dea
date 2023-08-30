package repository

import (
	"api/db"
	"api/models"
	"api/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func GetPersonById(id int64) (person models.Person, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	conn.First(&person, id)

	if person.IdPerson == 0 {
		person = models.Person{}
	}

	return
}

func PostPerson(person models.Person) (models.Person, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return person, err
	}

	passwordEncrypted, salt := utils.GenerateEncryptedPassword(person.Password)
	person.Password = passwordEncrypted
	person.Salt = salt

	conn.Create(&person)

	return GetPersonById(person.IdPerson)
}

func VerifyPersonByDocument(docNumber string) (found bool) {
	conn, err := db.OpenConnection()
	if err != nil {
		return false
	}

	var person models.Person
	if err := conn.Where("numeroDocumento = ?", docNumber).First(&person).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}

	return true
}

func GetAllPerson() (persons []models.Person, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	conn.Find(&persons)

	return
}

func PutPerson(person models.Person) (err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	passwordEncrypted, salt := utils.GenerateEncryptedPassword(person.Password)
	person.Password = passwordEncrypted
	person.Salt = salt

	result := conn.Where("idpessoa = ?", person.IdPerson).Updates(person)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("nenhum dado foi atualizado")
	}

	return nil
}

func DeletePersonById(id int64) error {
	conn, err := db.OpenConnection()

	if err != nil {
		return err
	}

	result := conn.Delete(&models.Person{}, id)

	return result.Error
}

func GetPersonByDocNumber(docNumber string) (person models.Person, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return person, err
	}

	conn.Where("numeroDocumento = ?", docNumber).First(&person)

	if person.IdPerson == 0 {
		person = models.Person{}
	}

	return person, nil
}
