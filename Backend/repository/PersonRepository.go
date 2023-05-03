package repository

import (
	"api/db"
	"api/models"
	"log"
)

func Get(id int64) (pessoa models.Person, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	row := conn.First(&pessoa, id)
	log.Printf("row: %v", row)

	return
}
