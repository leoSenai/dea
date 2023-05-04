package repository

import (
	"api/db"
	"api/models"
)

func Get(id int64) (pessoa models.Person, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	conn.First(&pessoa, id)

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

func VerifyPersonByDocument(docNumber string) (found bool) {
	conn, err := db.OpenConnection()
	if err != nil {
		return false
	}

	if err := conn.Where("numeroDocumento = ?", docNumber).First(&models.Person{}).Error; err != nil {
		return false
	}
	return true
}
