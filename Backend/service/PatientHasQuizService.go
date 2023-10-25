package service

import (
	"api/models"
	"api/repository"
	"time"
)

func GetPatientQuizByQuizID(idQuiz int64) (patientHasQuizz []models.PatientHasQuiz, err error) {
	patientHasQuizz, err = repository.GetPatientQuizByQuizID(idQuiz)
	return patientHasQuizz, err
}

func GetPatientQuizByPatientID(idPatient int64) (patientHasQuizz []models.PatientHasQuiz, err error) {
	patientHasQuizz, err = repository.GetPatientQuizByPatientID(idPatient)
	return patientHasQuizz, err
}

func PostPatientQuiz(patientHasQuiz models.PatientHasQuiz) (err error) {

	patientHasQuiz.AnsweredIn = time.Now().Local().String()

	patientHasQuiz, err = repository.PostPatientQuiz(patientHasQuiz)

	return nil
}

func PutPatientQuiz(patientHasQuiz models.PatientHasQuiz) (err error) {

	patientHasQuiz.AnsweredIn = time.Now().Local().String()
	err = repository.PutPatientQuiz(patientHasQuiz)
	return err
}
