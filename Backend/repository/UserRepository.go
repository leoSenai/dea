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

func GetAllUser() (users []models.User, err error) {
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
	conn.First(&userBack, userPost.IdUser)
	return
}

func PutUser(userPut models.User) (userBack models.User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	userBack = userPut
	userBack.IdUser = 0

	if userPut.IdUser != 0 {
		row := conn.Table("usuario").Where("idusuario = ?", userPut.IdUser).Updates(&userBack)
		log.Printf("row: %v", row)
		conn.First(&userBack, userPut.IdUser)
	}

	return
}
