package service

import (
	"api/models"
	"api/repository"
)

func GetQuizById(id int64) (quiz models.Quiz, err error) {
	quiz, err = repository.GetQuizById(int64(id))
	return quiz, err
}

func GetAllQuiz() (quiz []models.Quiz, err error) {
	quizs, err := repository.GetAllQuiz()
	return quizs, err
}

func PostQuiz(quizPost models.Quiz) (quizBack models.Quiz, err error) {
	quizBack, err = repository.PostQuiz(quizPost)
	return quizBack, err
}

func PutQuiz(quizPut models.Quiz) (quizBack models.Quiz, err error) {
	quizBack, err = repository.PutQuiz(quizPut)
	return quizBack, err
}
