package repository

import (
	"api/db"
	"api/models"
	"log"
)

func GetUserById(id int64) (user models.User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	row := conn.First(&user, id)
	log.Printf("row: %v", row)

	return
}

func GetAllUsers() (users []models.User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	rows := conn.Find(&users)
	log.Printf("rows: %v", rows)

	return
}

func PostUser(userPost models.User) (userBack models.User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	row := conn.Create(&userPost)
	log.Printf("row: %v", row)

	return
}
