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

	result := conn.Create(&quizPost)

	return
}

func PutQuiz(quizPut models.Quiz) (err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	conn.Table("questionario").Where("idquestionario = ?", quizPut.IdQuiz).Updates(&quizPut)

	return
}
