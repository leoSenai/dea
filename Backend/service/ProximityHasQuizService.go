package service

import (
	"api/models"
	"api/models/dtos"
	"api/repository"
	"fmt"
	"time"
)

func GetProximityQuizByQuizPatientPersonIDs(proximityHasQuizDto dtos.ProximityHasQuizDto) (proximityHasQuiz models.ProximityHasQuiz, err error) {
	proximityHasQuiz, err = repository.GetProximityQuizByQuizPatientPersonIDs(proximityHasQuizDto.IdQuiz, proximityHasQuizDto.ProximityIdPatient, proximityHasQuizDto.ProximityIdPerson)
	return proximityHasQuiz, err
}

func GetProximityQuizByQuizID(idQuiz int64) (proximityHasQuizz []models.ProximityHasQuiz, err error) {
	proximityHasQuizz, err = repository.GetProximityQuizByQuizID(idQuiz)
	return proximityHasQuizz, err
}

func GetProximityQuizByPatientID(idPatient int64) (proximityHasQuizz []models.ProximityHasQuiz, err error) {
	proximityHasQuizz, err = repository.GetProximityQuizByPatientID(idPatient)
	return proximityHasQuizz, err
}

func GetProximityQuizByPersonID(idPerson int64) (proximityHasQuizz []models.ProximityHasQuiz, err error) {
	proximityHasQuizz, err = repository.GetProximityQuizByPersonID(idPerson)
	return proximityHasQuizz, err
}

func PostProximityQuiz(proximityHasQuiz models.ProximityHasQuiz) (err error) {
	existingProximityHasQuiz, err := repository.GetProximityQuizByQuizPatientPersonIDs(proximityHasQuiz.IdQuiz, proximityHasQuiz.ProximityIdPatient, proximityHasQuiz.ProximityIdPerson)
	if err != nil {
		return err
	}

	if existingProximityHasQuiz.IdQuiz != 0 {
		return fmt.Errorf("Já esta relacionado no banco de dados")
	}

	proximityHasQuiz.AnsweredIn = time.Now()

	proximityHasQuiz, err = repository.PostProximityQuiz(proximityHasQuiz)

	return nil
}

func PutProximityQuiz(proximityHasQuiz models.ProximityHasQuiz) (err error) {
	existingProximityHasQuiz, err := repository.GetProximityQuizByQuizPatientPersonIDs(proximityHasQuiz.IdQuiz, proximityHasQuiz.ProximityIdPatient, proximityHasQuiz.ProximityIdPerson)
	if err != nil {
		return err
	}

	if existingProximityHasQuiz.IdQuiz != 0 {
		proximityHasQuiz.AnsweredIn = time.Now()
		err := repository.PutProximityQuiz(proximityHasQuiz)
		return err
	} else {
		return fmt.Errorf("Relacionamento não cadastrado no banco de dados")
	}
}
