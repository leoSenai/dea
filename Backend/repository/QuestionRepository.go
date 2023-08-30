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

func PutQuestionsBulk(questions []models.Question) (insertedQuestions []models.Question, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}

	for _, q := range questions {
		var insertedQuestion models.Question
		if q.IdQuestion == 0 {
			row := conn.Create(&q)
			if row.Error != nil {
				// log.Printf("Error inserting question: %s", row.Error.Error())
				continue
			}
			insertedQuestion = q
		} else {
			row := conn.Model(&q).Updates(map[string]interface{}{
				"desc":   q.Desc,
				"idquiz": q.IdQuiz,
			})
			if row.Error != nil {
				// log.Printf("Error updating question: %s", row.Error.Error())
				continue
			}
			insertedQuestion = q
		}
		insertedQuestions = append(insertedQuestions, insertedQuestion)
	}
	return insertedQuestions, nil
}

func GetQuestionsByQuiz(idQuiz int) (questions []models.Question, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return nil, err
	}

	rows := conn.Where("idquestionario = ?", idQuiz).Find(&questions)
	if rows.Error != nil {
		// log.Printf("Error fetching questions by quiz: %s", rows.Error.Error())
		return nil, rows.Error
	}

	return questions, nil
}

func DeleteQuestionById(id int64) error {
	conn, err := db.OpenConnection()
	if err != nil {
		return err
	}

	row := conn.Delete(&models.Question{}, id)
	if row.Error != nil {
		log.Printf("Erro ao deletar questão: %s", row.Error.Error())
		return row.Error
	}

	return nil
}
