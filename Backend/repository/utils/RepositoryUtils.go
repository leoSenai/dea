package utils

import (
	"api/db"
	"api/models"
)

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
