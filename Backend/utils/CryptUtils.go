package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSha256(password string) (sha256password string) {

	sha256o := sha256.New()
	sha256o.Write([]byte(password))
	sha256password = hex.EncodeToString(sha256o.Sum(nil))

	return sha256password
}

func GenerateSalt() (salt string) {

	saltByte := make([]byte, 15, 20)
	rand.Read(saltByte)
	salt = hex.EncodeToString(saltByte)

	return salt
}
