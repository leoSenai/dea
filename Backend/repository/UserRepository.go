package repository

import (
	"api/db"
	"api/models"
	generalUtils "api/utils"
	"log"
)

func GetUserById(id int64) (user models.User, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.First(&user, id)
	log.Printf("row: %v", row)

	return
}

func GetUserByLogin(login string) (user models.User, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.First(&user, "email = ? OR telefone = ?", login, login)
	log.Printf("row: %v", row)

	return
}

func GetAllUser() (users []models.User, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	rows := conn.Find(&users)
	log.Printf("rows: %v", rows)

	return
}

func PostUser(userPost models.User) (userBack models.User, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	passwordEncrypted, salt := generalUtils.GenerateEncryptedPassword(userPost.Password)
	userPost.Password = passwordEncrypted
	userPost.Salt = salt
	userPost.IdServices = 1

	row := conn.Create(&userPost)
	log.Printf("row: %v", row)

	return
}

func PutUser(userPut models.User) (userBack models.User, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	if userPut.Password != "" {
		passwordEncrypted, salt := generalUtils.GenerateEncryptedPassword(userPut.Password)
		userPut.Password = passwordEncrypted
		userPut.Salt = salt
		userPut.IdServices = 1
	}

	if userPut.IdUser != 0 {

		cbo, _ := GetCboByCode(int64(userPut.IdCbo))
		userPut.IdCbo = cbo.IdCbo

		row := conn.Table("usuario").Where("idusuario = ?", userPut.IdUser).Updates(&userPut)
		conn.Table("usuario").Select("ativo").Where("idusuario = ?", userPut.IdUser).Updates(&userPut)
		log.Printf("row: %v", row)
	}

	return
}

func VerifyUserExistanceByDocument(email string) (exist bool) {

	conn, err := db.GetDB()
	if err != nil {
		return
	}

	var userFound models.User

	conn.First(&userFound, "email = ?", email)

	if userFound.IdUser != 0 {
		return true
	}

	return false
}
