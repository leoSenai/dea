package models

import (
	"api/db"
	"log"
)

func Get(id int64) (pessoa Pessoa, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	row := conn.First(&pessoa, id)
	log.Printf("row: %v", row)

	return
}

func Hello() {
	log.Printf("Hello Word")
}
