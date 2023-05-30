package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSha256(password string) (sha256password string) {

	sha256o := sha256.New()
	sha256o.Write([]byte(password))
	sha256password = hex.EncodeToString(sha256o.Sum(nil))

	return sha256password
}
