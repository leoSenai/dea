package service

import (
	"api/models"
	"api/repository"
	"fmt"
)

func GetAskingById(id int64) (asking models.Asking, err error) {
	asking, err = repository.GetAskingById(int64(id))
	return asking, err
}

func PostAsking(asking models.Asking) (err error) {
	found := repository.VerifyAskingByDescription(asking.Desc)

	if found {
		return fmt.Errorf("Pergunta com a DESCRIÇÃO %s já cadastrado!", asking.Desc)
	} else {
		err = repository.PostAsking(asking)
	}

	return err
}

func GetAllAsking() (askings []models.Asking, err error) {
	askings, err = repository.GetAllAsking()
	return askings, err
}

func PutAsking(askingUpdate models.Asking) (err error) {
	var asking models.Asking

	asking, err = repository.GetAskingById(askingUpdate.IdAsking)
	if err != nil {
		return err
	}

	if asking.IdAsking == 0 {
		return fmt.Errorf("Pergunta não está cadastrado no sistema!")
	}

	err = repository.PutAsking(askingUpdate)

	return
}
