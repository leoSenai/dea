package repository

import (
	"api/db"
	"api/models"
	"log"
)

func GetPersonById(id int64) (person models.Person, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	row := conn.First(&person, id)
	log.Printf("row: %v", row)

	return
}
