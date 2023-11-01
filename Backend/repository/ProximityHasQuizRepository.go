package repository

import (
	"api/db"
	"api/models"
	"fmt"
	"strings"
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

	if err := conn.Where("questionario_idquestionario = ? AND finalizado = 0", quizID).Find(&proximityHasQuizzes).Error; err != nil {
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

	var proximity models.Proximity

	if err != nil {
		return proximityHasQuiz, err
	}

	conn.First(&proximity, "pessoa_idpessoa = ?", proximityHasQuiz.ProximityIdPerson)

	proximityHasQuiz = models.ProximityHasQuiz{
		IdQuiz:             proximityHasQuiz.IdQuiz,
		ProximityIdPatient: proximity.IdPatient,
		ProximityIdPerson:  proximityHasQuiz.ProximityIdPerson,
		Finished:           0,
		Answers:            "",
	}
	conn.Omit("respondido_em", "respostas").Create(&proximityHasQuiz)

	return proximityHasQuiz, err
}

func PutProximityQuiz(proximityHasQuiz models.ProximityHasQuiz) error {
	conn, err := db.GetDB()

	if err != nil {
		return err
	}

	proximityHasQuiz.AnsweredIn = strings.Split(proximityHasQuiz.AnsweredIn, ".")[0]

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

func GetProximityQuizByQuizPersonID(idQuiz int64, idPerson int64) (alreadyExist bool, proximityHasQuizzes []models.ProximityHasQuiz, err error) {
	conn, err := db.GetDB()

	var count int64
	row := conn.Table("proximidade_has_questionario").Where("proximidade_pessoa_idpessoa = ? AND questionario_idquestionario = ?", idPerson, idQuiz).Find(&proximityHasQuizzes)
	if row.Error != nil {
		err = row.Error
		return
	}
	row.Count(&count)
	if count > 0 {
		alreadyExist = true
	} else {
		alreadyExist = false
	}

	return
}

func DeleteProximityHasQuiz(proximityHasQuizRemoved models.ProximityHasQuiz) (err error) {
	conn, err := db.GetDB()

	conn.Omit("respondido_em", "respostas").Where("proximidade_paciente_idpaciente = ? AND proximidade_pessoa_idpessoa = ? AND questionario_idquestionario = ? AND finalizado = 0", proximityHasQuizRemoved.ProximityIdPatient, proximityHasQuizRemoved.ProximityIdPerson, proximityHasQuizRemoved.IdQuiz).Delete(&proximityHasQuizRemoved)
	if err != nil {
		return err
	}

	return
}
