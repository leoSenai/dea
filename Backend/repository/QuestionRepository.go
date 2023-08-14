package repository

import (
	"api/db"
	"api/models"
	"log"
)

func GetQuestionById(id int64) (question models.Question, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	row := conn.First(&question, id)
	log.Printf("row: %v", row)

	return
}

func GetAllQuestion() (questions []models.Question, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	rows := conn.Find(&questions)
	log.Printf("rows: %v", rows)

	return
}

func PostQuestion(questionPost models.Question) (questionBack models.Question, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	row := conn.Create(&questionPost)
	log.Printf("row: %v", row)
	conn.First(&questionBack, questionPost.IdQuestion)

	return
}

func PutQuestion(questionPut models.Question) (questionBack models.Question, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	questionBack = questionPut
	questionBack.IdQuestion = 0

	if questionPut.IdQuestion != 0 {
		row := conn.Table("questao").Where("idquestao = ?", questionPut.IdQuestion).Updates(&questionBack)
		log.Printf("row: %v", row)
		conn.First(&questionBack, questionPut.IdQuestion)
	}

	return
}
