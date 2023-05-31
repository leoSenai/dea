package utils

import (
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

func generateSalt() (salt string) {

	saltByte := make([]byte, 15, 20)
	rand.Read(saltByte)
	salt = hex.EncodeToString(saltByte)

	return salt
}
