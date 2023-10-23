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

func GenerateRandomPassword(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // Caracteres permitidos
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	for i := 0; i < length; i++ {
		randomBytes[i] = charset[int(randomBytes[i])%len(charset)]
	}

	return string(randomBytes), nil
}

func IsPasswordValid(login string, password string) (isValid bool) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	var userVerify models.User
	userVerify.Password = password
	userVerify.Name = login

	var user models.User
	conn.First(&user, "email = ? OR telefone = ?", userVerify.Name, userVerify.Name)

	sha256o := sha256.New()
	passwordSalted := userVerify.Password + user.Salt

	sha256o.Write([]byte(passwordSalted))
	encryptedPassword := hex.EncodeToString(sha256o.Sum(nil))

	print("Teste: 1 " + user.Password)
	print("Teste: 2 " + encryptedPassword)
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
