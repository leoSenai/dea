package utils

import (
	"api/db"
	"api/models"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateEncryptedPassword(password string) (encryptedPassword string, salt string) {

	sha256o := sha256.New()

	salt = generateSalt()
	passwordSalted := password + salt

	sha256o.Write([]byte(passwordSalted))
	encryptedPassword = hex.EncodeToString(sha256o.Sum(nil))

	return encryptedPassword, salt
}

func IsPasswordValid(login string, password string) (isValid bool) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	var userVerify models.User
	userVerify.Password = password
	userVerify.Name = login

	var user models.User
	conn.First(&user, "nome = ?", userVerify.Name)

	sha256o := sha256.New()
	passwordSalted := userVerify.Password + user.Salt

	sha256o.Write([]byte(passwordSalted))
	encryptedPassword := hex.EncodeToString(sha256o.Sum(nil))

	if encryptedPassword == user.Password {
		return true
	} else {
		return false
	}

}

func generateSalt() (salt string) {

	saltByte := make([]byte, 15, 20)
	rand.Read(saltByte)
	salt = hex.EncodeToString(saltByte)

	return salt
}
