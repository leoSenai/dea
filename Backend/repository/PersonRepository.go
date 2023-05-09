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

func PostPerson(person models.Person) (err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	conn.Create(person)

	return
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

	result := conn.Where("idpessoa = ?", person.IdPerson).Updates(models.Person{Name: person.Name, BornDate: person.BornDate, DocNumber: person.DocNumber, DocType: person.DocType, Password: person.Password, Salt: person.Salt})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("Nenhum dado foi atualizado")
	}

	return nil
}
