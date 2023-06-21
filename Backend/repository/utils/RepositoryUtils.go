package utils

import (
	"api/db"
	"api/models"
)

func VerifyUserExistanceByDocument(idCbo int) (exist bool) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	var userFound models.User

	conn.First(&userFound, "cbo_idcbo = ?", idCbo)

	if userFound.IdUser != 0 {
		return true
	}

	return false
}
