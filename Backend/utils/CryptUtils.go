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

func IsPasswordValid(login string, password string) bool {
	conn, err := db.GetDB()
	if err != nil {
		return false
	}

	var userDynamic UserDynamic
	var userFound, personFound, patientFound bool

	var userVerify models.User
	userVerify.Password = password
	userVerify.Name = login

	var user models.User
	var person models.Person
	var patient models.Patient

	userFound = conn.First(&user, "email = ? OR telefone = ?", userVerify.Name, userVerify.Name).Error == nil
	if !userFound {
		personFound = conn.First(&person, "email = ? OR telefone = ?", userVerify.Name, userVerify.Name).Error == nil
		if !personFound {
			patientFound = conn.First(&patient, "email = ? OR telefone = ?", userVerify.Name, userVerify.Name).Error == nil
			if !patientFound {
				return false
			}
			userDynamic = UserDynamic{
				Email:    patient.Email,
				Phone:    patient.Phone,
				Password: patient.Password,
				Salt:     patient.Salt,
			}
		} else {
			userDynamic = UserDynamic{
				Email:    person.Email,
				Phone:    person.Phone,
				Password: person.Password,
				Salt:     person.Salt,
			}
		}
	} else {
		userDynamic = UserDynamic{
			Email:    user.Email,
			Phone:    user.Phone,
			Password: user.Password,
			Salt:     user.Salt,
		}
	}

	sha256o := sha256.New()
	passwordSalted := userVerify.Password + userDynamic.Salt

	sha256o.Write([]byte(passwordSalted))
	encryptedPassword := hex.EncodeToString(sha256o.Sum(nil))

	return encryptedPassword == userDynamic.Password
}

func generateSalt() (salt string) {

	saltByte := make([]byte, 15, 20)
	rand.Read(saltByte)
	salt = hex.EncodeToString(saltByte)

	return salt
}

type UserDynamic struct {
	Email    string
	Phone    string
	Password string
	Salt     string
}
