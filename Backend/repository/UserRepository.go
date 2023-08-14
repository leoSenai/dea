package repository

import (
	"api/db"
	"api/models"
	repositoryUtils "api/repository/utils"
	generalUtils "api/utils"
	"errors"
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

	found := repositoryUtils.VerifyUserExistanceByDocument(userPost.Email)

	if found {
		err = errors.New("Esse usuário já foi cadastrado no sistema!")
	} else {

		passwordEncrypted, salt := generalUtils.GenerateEncryptedPassword(userPost.Password)
		userPost.Password = passwordEncrypted
		userPost.Salt = salt

		row := conn.Create(&userPost)
		log.Printf("row: %v", row)

	}

	return
}

func PutUser(userPut models.User) (userBack models.User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	if userPut.Password != "" {
		passwordEncrypted, salt := generalUtils.GenerateEncryptedPassword(userPut.Password)
		userPut.Password = passwordEncrypted
		userPut.Salt = salt
	}

	if userPut.IdUser != 0 {
		row := conn.Table("usuario").Where("idusuario = ?", userPut.IdUser).Updates(&userPut)
		log.Printf("row: %v", row)
	}

	return
}
