package repository

import (
	"api/db"
	"api/models"
	"log"
)

func GetQuizById(id int64) (quiz models.Quiz, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.First(&quiz, id)
	log.Printf("row: %v", row)

	return
}

func GetAllQuiz() (quizs []models.Quiz, quizsFinisheds []models.QuizFinished, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	rows := conn.Find(&quizs)
	log.Printf("rows: %v", rows)

	var patientHasQuizzes []models.PatientHasQuiz

	rows = conn.Find(&patientHasQuizzes)
	log.Printf("rows: %v", rows)

	var proximityHasQuizzes []models.ProximityHasQuiz

	rows = conn.Find(&proximityHasQuizzes)
	log.Printf("rows: %v", rows)

	for _, patientHasQuiz := range patientHasQuizzes {
		if patientHasQuiz.Finished == 1 {
			quizsFinisheds = append(quizsFinisheds, models.QuizFinished{
				IdQuiz:   patientHasQuiz.IdQuiz,
				Finished: patientHasQuiz.Finished,
			})
		}
	}

	for _, proximityHasQuiz := range proximityHasQuizzes {
		if proximityHasQuiz.Finished == 1 {
			quizsFinisheds = append(quizsFinisheds, models.QuizFinished{
				IdQuiz:   proximityHasQuiz.IdQuiz,
				Finished: proximityHasQuiz.Finished,
			})
		}
	}

	return quizs, quizsFinisheds, err
}

func PostQuiz(quizPost models.Quiz) (quizBack models.Quiz, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	conn.Create(&quizPost)
	quizBack = quizPost
	return
}

func PutQuiz(quizPut models.Quiz) (err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	conn.Table("questionario").Where("idquestionario = ?", quizPut.IdQuiz).Updates(&quizPut)

	return
}
