package service

import (
	"api/models"
	"api/repository"
)

func GetAnamneseById(id int64) (anamnese models.Anamnese, err error) {
	anamnese, err = repository.GetAnamneseById(int64(id))
	return anamnese, err
}

func GetAllAnamnese() (anamnese []models.Anamnese, err error) {
	anamneses, err := repository.GetAllAnamnese()
	return anamneses, err
}

func PostAnamnese(anamnesePost models.Anamnese) (anamneseBack models.Anamnese, err error) {
	anamneseBack, err = repository.PostAnamnese(anamnesePost)
	return anamneseBack, err
}

func PutAnamnese(anamnesePut models.Anamnese) (anamneseBack models.Anamnese, err error) {
	anamneseBack, err = repository.PutAnamnese(anamnesePut)
	return anamneseBack, err
}

func GetAnamneseByIdAndPatientId(userId int64, patientId int64) (anamnese models.Anamnese, err error) {
	anamnese, err = repository.GetAnamneseByIdAndPatientId(int64(userId), int64(patientId))
	return anamnese, err
}
