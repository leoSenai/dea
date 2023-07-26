package service

import (
	"api/models"
	"api/repository"
	"fmt"
	"time"
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
	now := time.Now().Format("2006-01-02 15:04:05")
	quizPost.Created = now
	quizPost.Updated = now
	quizBack, err = repository.PostQuiz(quizPost)
	return quizBack, err
}

func PutQuiz(quizPut models.Quiz) (err error) {
	found, err := repository.VerifyResponseQuizById(quizPut.IdQuiz)
	if err != nil {
		return fmt.Errorf("Erro ao verificar questionário: %v", err)
	}

	if found {
		return fmt.Errorf("Questionário já respondido, não é possível atualizar!")
	}

	err = repository.PutQuiz(quizPut)
	if err != nil {
		return fmt.Errorf("Erro ao atualizar o questionário: %v", err)
	}

	return nil
}
