package service

import (
	"api/models"
	"api/repository"
)

func GetAnamneseHasAskingByAnamneseId(id int64) (anamneseHasAsking []models.AnamneseHasAsking, err error) {
	anamneseHasAsking, err = repository.GetAnamneseHasAskingByAnamneseId(int64(id))
	return anamneseHasAsking, err
}

func GetAnamneseHasAskingByAskingId(id int64) (anamneseHasAsking []models.AnamneseHasAsking, err error) {
	anamneseHasAsking, err = repository.GetAnamneseHasAskingByAskingId(int64(id))
	return anamneseHasAsking, err
}

func GetAllAnamneseHasAsking() (anamneseHasAskings []models.AnamneseHasAsking, err error) {
	anamneseHasAskings, err = repository.GetAllAnamneseHasAsking()
	return anamneseHasAskings, err
}

func PostAnamneseHasAsking(anamneseHasAskingPost models.AnamneseHasAsking) (anamneseHasAskingBack models.AnamneseHasAsking, err error) {
	anamneseHasAskingBack, err = repository.PostAnamneseHasAsking(anamneseHasAskingPost)
	return anamneseHasAskingBack, err
}

func DeleteAnamneseHasAsking(anamneseHasAskingDelete models.AnamneseHasAsking) (anamneseHasAskingBack models.AnamneseHasAsking, err error) {
	anamneseHasAskingBack, err = repository.DeleteAnamneseHasAsking(anamneseHasAskingDelete)
	return anamneseHasAskingBack, err
}
