package service

import (
	"api/models"
	"api/models/dtos"
	"api/repository"
)

func GetUserById(id int64) (user dtos.UserDTO, err error) {
	user, err = repository.GetUserById(int64(id))
	return user, err
}

func GetAllUser() (user []dtos.UserDTO, err error) {
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
