package service

import (
	"api/models"
	"api/models/dtos"
	"api/repository"
	"api/utils"
	"fmt"
)

func GetUserById(id int64) (user models.User, err error) {
	user, err = repository.GetUserById(int64(id))
	return user, err
}

func GetUserByLogin(login string) (user models.User, err error) {
	user, err = repository.GetUserByLogin(login)
	return user, err
}

func GetAllUser() (user []models.User, err error) {
	users, err := repository.GetAllUser()
	return users, err
}

func PostUser(userPost models.User) (userBack models.User, err error) {
	userBack, err = repository.PostUser(userPost)
	return userBack, err
}

func PutUser(userPut models.User) (userBack models.User, err error) {
	userBack, err = repository.PutUser(userPut)
	return userBack, err
}

func ResetPasswordUser(userDto dtos.UserDTO) (isReseted bool, err error) {
	fmt.Println(userDto.IdUser)
	user, err := repository.GetUserById(int64(userDto.IdUser))
	to := user.Email
	if err != nil {
		return false, err
	}

	password, err := utils.GenerateRandomPassword(8)

	if err != nil {
		return false, fmt.Errorf("erro ao gerar senha")
	}

	user.Password = password

	user, err = repository.PutUser(user)

	if err != nil {
		return false, err
	}

	subject := "DEA - Nova senha de acesso"
	body := "Sua nova senha: " + password + "\n\n"
	fmt.Println(to)
	fmt.Println(user.Name)
	fmt.Println(password)

	err = utils.SendEmail(to, subject, body)
	if err != nil {
		return false, fmt.Errorf("erro ao enviar e-mail de reset de senha do usuario")
	}

	return true, nil
}
