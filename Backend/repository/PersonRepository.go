package repository

import (
	"api/db"
	"api/models"
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

func InsertPerson(person models.Person) (err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	conn.Create(person)

	return
}

func VerifyPersonByDocument(docNumber string) (found bool, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return false, err
	}

	var person models.Person
	if err := conn.Where("numeroDocumento = ?", docNumber).First(&person).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func GetAllPersons() (persons []models.Person, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	conn.Find(&persons)

	return
}

func UpdatePerson(person models.Person) (err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	result := conn.Where("idpessoa = ?", person.IdPerson).Updates(models.Person{Name: person.Name, BornDate: person.BornDate, DocNumber: person.DocNumber, DocType: person.DocType, Password: person.Password, Salt: person.Salt})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("Nenhum dado foi atualizado")
	}

	return nil
}
