package service

import (
	"api/models"
	"api/repository"
)

func GetUserById(id int64) (user models.User, err error) {
	user, err = repository.GetUserById(int64(id))
	return user, err
}

func GetAllUsers() (user []models.User, err error) {
	users, err := repository.GetAllUsers()
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
