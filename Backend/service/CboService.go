package service

import (
	"api/models"
	"api/repository"
)

func GetCboById(id int64) (cbo models.Cbo, err error) {
	cbo, err = repository.GetCboById(int64(id))
	return cbo, err
}

func GetAllCbos() (cbos []models.Cbo, err error) {
	cbos, err = repository.GetAllCbos()
	return cbos, err
}

func PostCbo(cboPost models.Cbo) (cboBack models.Cbo, err error) {
	cboBack, err = repository.PostCbo(cboPost)
	return cboBack, err
}

func PutCbo(cboPut models.Cbo) (cboBack models.Cbo, err error) {
	cboBack, err = repository.PutCbo(cboPut)
	return cboBack, err
}
