package service

import (
	"api/models"
	"api/repository"
)

func GetQuestionById(id int64) (question models.Question, err error) {
	question, err = repository.GetQuestionById(int64(id))
	return question, err
}

func GetAllQuestion() (question []models.Question, err error) {
	questions, err := repository.GetAllQuestion()
	return questions, err
}

func PostQuestion(questionPost models.Question) (questionBack models.Question, err error) {
	questionBack, err = repository.PostQuestion(questionPost)
	return questionBack, err
}

func PutQuestion(questionPut models.Question) (questionBack models.Question, err error) {
	questionBack, err = repository.PutQuestion(questionPut)
	return questionBack, err
}
