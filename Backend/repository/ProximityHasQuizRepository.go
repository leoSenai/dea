package repository

import (
	"api/db"
	"api/models"
	"fmt"
)

func GetProximityQuizByQuizPatientPersonIDs(idQuiz int64, idPatient int64, idPerson int64) (proximityHasQuiz models.ProximityHasQuiz, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return proximityHasQuiz, err
	}

	if err := conn.Where("proximidade_paciente_idpaciente = ? AND proximidade_pessoa_idpessoa = ? AND questionario_idquestionario = ?", idPatient, idPerson, idQuiz).First(&proximityHasQuiz).Error; err != nil {
		return proximityHasQuiz, err
	}

	return proximityHasQuiz, nil
}

func GetProximityQuizByQuizID(quizID int64) ([]models.ProximityHasQuiz, error) {
	conn, err := db.GetDB()
	if err != nil {
		return nil, err
	}

	var proximityHasQuizzes []models.ProximityHasQuiz

	if err := conn.Where("questionario_idquestionario = ?", quizID).Find(&proximityHasQuizzes).Error; err != nil {
		return nil, err
	}

	return proximityHasQuizzes, nil
}

func GetProximityQuizByPatientID(patientID int64) ([]models.ProximityHasQuiz, error) {
	conn, err := db.GetDB()
	if err != nil {
		return nil, err
	}

	var proximityHasQuizzes []models.ProximityHasQuiz

	if err := conn.Where("proximidade_paciente_idpaciente = ?", patientID).Find(&proximityHasQuizzes).Error; err != nil {
		return nil, err
	}

	return proximityHasQuizzes, nil
}

func GetProximityQuizByPersonID(personID int64) ([]models.ProximityHasQuiz, error) {
	conn, err := db.GetDB()
	if err != nil {
		return nil, err
	}

	var proximityHasQuizzes []models.ProximityHasQuiz

	if err := conn.Where("proximidade_pessoa_idpessoa = ?", personID).Find(&proximityHasQuizzes).Error; err != nil {
		return nil, err
	}

	return proximityHasQuizzes, nil
}

func PostProximityQuiz(proximityHasQuiz models.ProximityHasQuiz) (models.ProximityHasQuiz, error) {
	conn, err := db.GetDB()

	if err != nil {
		return proximityHasQuiz, err
	}

	conn.Create(&proximityHasQuiz)

	return proximityHasQuiz, err
}

func PutProximityQuiz(proximityHasQuiz models.ProximityHasQuiz) error {
	conn, err := db.GetDB()

	if err != nil {
		return err
	}

	result := conn.Where("proximidade_paciente_idpaciente = ? AND proximidade_pessoa_idpessoa = ? AND questionario_idquestionario = ?", proximityHasQuiz.ProximityIdPatient, proximityHasQuiz.ProximityIdPerson, proximityHasQuiz.IdQuiz).Updates(
		models.ProximityHasQuiz{
			Finished:   proximityHasQuiz.Finished,
			Answers:    proximityHasQuiz.Answers,
			AnsweredIn: proximityHasQuiz.AnsweredIn,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("nenhum dado foi atualizado")
	}

	return nil
}
