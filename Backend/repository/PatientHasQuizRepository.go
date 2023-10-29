package repository

import (
	"api/db"
	"api/models"
	"fmt"
)

func VerifyResponseQuizById(idQuiz int64) (hasQuiz bool, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	var count int64
	row := conn.Table("paciente_has_questionario").Where("respondido_em IS NOT NULL AND questionario_idquestionario = ?", idQuiz).Count(&count)
	if row.Error != nil {
		err = row.Error
		return
	}

	if count > 0 {
		hasQuiz = true
	}

	return
}

func GetPatientQuizByQuizID(quizID int64) ([]models.PatientHasQuiz, error) {
	conn, err := db.GetDB()
	if err != nil {
		return nil, err
	}

	var patientHasQuizzes []models.PatientHasQuiz

	if err := conn.Where("questionario_idquestionario = ?", quizID).Find(&patientHasQuizzes).Error; err != nil {
		return nil, err
	}

	return patientHasQuizzes, nil
}

func GetPatientQuizByPatientID(patientID int64) ([]models.PatientHasQuiz, error) {
	conn, err := db.GetDB()
	if err != nil {
		return nil, err
	}

	var patientHasQuizzes []models.PatientHasQuiz

	if err := conn.Where("proximidade_paciente_idpaciente = ?", patientID).Find(&patientHasQuizzes).Error; err != nil {
		return nil, err
	}

	return patientHasQuizzes, nil
}

func PostPatientQuiz(patientHasQuiz models.PatientHasQuiz) (models.PatientHasQuiz, error) {
	conn, err := db.GetDB()

	if err != nil {
		return patientHasQuiz, err
	}

	conn.Create(&patientHasQuiz)

	return patientHasQuiz, err
}

func PutPatientQuiz(patientHasQuiz models.PatientHasQuiz) error {
	conn, err := db.GetDB()

	if err != nil {
		return err
	}

	result := conn.Where("proximidade_paciente_idpaciente = ? AND proximidade_pessoa_idpessoa = ? AND questionario_idquestionario = ?", patientHasQuiz.IdPatient, patientHasQuiz.IdQuiz).Updates(
		models.PatientHasQuiz{
			Finished:   patientHasQuiz.Finished,
			Answers:    patientHasQuiz.Answers,
			AnsweredIn: patientHasQuiz.AnsweredIn,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("nenhum dado foi atualizado")
	}

	return nil
}
