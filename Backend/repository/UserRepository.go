package repository

import (
	"api/db"
	"api/models"
	"log"
)

func GetQuizById(id int64) (quiz models.Quiz, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	row := conn.First(&quiz, id)
	log.Printf("row: %v", row)

	return
}

func GetAllQuiz() (quizs []models.Quiz, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	rows := conn.Find(&quizs)
	log.Printf("rows: %v", rows)

	return
}

func PostQuiz(quizPost models.Quiz) (quizBack models.Quiz, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	row := conn.Create(&quizPost)
	log.Printf("row: %v", row)
	conn.First(&quizBack, quizPost.IdQuiz)

	return
}

func PutQuiz(quizPut models.Quiz) (quizBack models.Quiz, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	quizBack = quizPut
	quizBack.IdQuiz = 0

	if quizPut.IdQuiz != 0 {
		row := conn.Table("questionario").Where("idquestionario = ?", quizPut.IdQuiz).Updates(&quizBack)
		log.Printf("row: %v", row)
		conn.First(&quizBack, quizPut.IdQuiz)
	}

	return
}
