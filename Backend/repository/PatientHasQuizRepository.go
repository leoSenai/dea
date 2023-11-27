package repository

import (
	"api/db"
	"api/models"
	"fmt"
	"strings"
)

func VerifyResponseQuizById(idQuiz int64) (hasQuiz bool, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	var countPatient int64
	row1 := conn.Table("paciente_has_questionario").Where("respondido_em IS NOT NULL AND questionario_idquestionario = ?", idQuiz).Count(&countPatient)
	if row1.Error != nil {
		err = row1.Error
		return
	}

	if countPatient > 0 {
		hasQuiz = true
	}

	var countPerson int64
	row2 := conn.Table("proximidade_has_questionario").Where("respondido_em IS NOT NULL AND questionario_idquestionario = ?", idQuiz).Count(&countPerson)
	if row2.Error != nil {
		err = row2.Error
		return
	}

	if countPerson > 0 {
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

	if err := conn.Where("questionario_idquestionario = ? AND finalizado = 0", quizID).Find(&patientHasQuizzes).Error; err != nil {
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

	if err := conn.Where("paciente_idpaciente = ?", patientID).Find(&patientHasQuizzes).Error; err != nil {
		return nil, err
	}

	return patientHasQuizzes, nil
}

func PostPatientQuiz(patientHasQuiz models.PatientHasQuiz) (models.PatientHasQuiz, error) {
	conn, err := db.GetDB()

	if err != nil {
		return patientHasQuiz, err
	}

	//patientHasQuiz.AnsweredIn = strings.Split(patientHasQuiz.AnsweredIn, ".")[0]

	conn.Create(&patientHasQuiz)

	return patientHasQuiz, err
}

func PutPatientQuiz(patientHasQuiz models.PatientHasQuiz) error {
	conn, err := db.GetDB()

	if err != nil {
		return err
	}

	patientHasQuiz.AnsweredIn = strings.Split(patientHasQuiz.AnsweredIn, ".")[0]

	result := conn.Where("paciente_idpaciente = ? AND questionario_idquestionario = ?", patientHasQuiz.IdPatient, patientHasQuiz.IdQuiz).Updates(
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

func GetPatientQuizByQuizPatientID(idQuiz int64, idPatient int64) (alreadyExist bool, patientHasQuizzes []models.PatientHasQuiz, err error) {
	conn, err := db.GetDB()

	var count int64
	row := conn.Table("paciente_has_questionario").Where("paciente_idpaciente = ? AND questionario_idquestionario = ?", idPatient, idQuiz).Find(&patientHasQuizzes)
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

	return alreadyExist, patientHasQuizzes, err
}

func DeletePatientHasQuiz(patientHasQuizRemoved models.PatientHasQuiz) (err error) {
	conn, err := db.GetDB()

	conn.Omit("respondido_em", "respostas").Where("paciente_idpaciente = ? AND questionario_idquestionario = ? AND finalizado = 0", patientHasQuizRemoved.IdPatient, patientHasQuizRemoved.IdQuiz).Delete(&patientHasQuizRemoved)
	if err != nil {
		return err
	}

	return
}
