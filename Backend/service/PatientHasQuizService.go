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

func GetPatientQuizByQuizPatientID(idQuiz int64, idPatient int64) (alreadyExist bool, patientHasQuizzes []models.PatientHasQuiz, err error) {
	alreadyExist, patientHasQuizzes, _ = repository.GetPatientQuizByQuizPatientID(idQuiz, idPatient)
	return alreadyExist, patientHasQuizzes, err
}

func GetPatientQuizByPatientID(idPatient int64) (patientHasQuizz []models.PatientHasQuiz, err error) {
	patientHasQuizz, err = repository.GetPatientQuizByPatientID(idPatient)
	return patientHasQuizz, err
}

func PostPatientQuiz(patientHasQuiz models.PatientHasQuiz) (err error) {

	//patientHasQuiz.AnsweredIn = time.Now().Local().String()

	patientHasQuiz, _ = repository.PostPatientQuiz(patientHasQuiz)

	return nil
}

func PutPatientQuiz(patientHasQuiz models.PatientHasQuiz) (err error) {

	patientHasQuiz.AnsweredIn = time.Now().Local().String()
	err = repository.PutPatientQuiz(patientHasQuiz)
	return err
}

func DeletePatientQuiz(patientHasQuizRemoved models.PatientHasQuiz) (err error) {
	err = repository.DeletePatientHasQuiz(patientHasQuizRemoved)
	if err != nil {
		return err
	}
	return
}
