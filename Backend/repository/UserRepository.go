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

	found := repositoryUtils.VerifyUserExistanceByDocument(userPost.Email)

	if found {
		err = errors.New("esse usuario ja foi cadastrado no sistema")
	} else {

		cbo, _ := GetCboByCode(int64(userPost.IdCbo))
		if cbo.IdCbo == 0 {
			err = errors.New("o codigo brasileiro de opcupacao informado nao existe na base de dados")
			return
		}

		passwordEncrypted, salt := generalUtils.GenerateEncryptedPassword(userPost.Password)
		userPost.Password = passwordEncrypted
		userPost.Salt = salt
		userPost.IdServices = 1
		userPost.IdCbo = cbo.IdCbo

		row := conn.Create(&userPost)
		log.Printf("row: %v", row)

	}

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
		log.Printf("row: %v", row)
	}

	return
}
