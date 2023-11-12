package controller

import (
	"api/models"
	"api/service"
	"api/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func GetQuizById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Cannot parse ID: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi especificado o id do questionário procurado.", "")

		return
	}

	quiz, err := service.GetQuizById(int64(id))
	if err != nil {
		log.Printf("Cannot find Get: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar questionário, houve um erro interno no servidor.", "")

		return
	} else if quiz.IdQuiz == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não foi possível encontrar o questionário.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Questionário encontrado com sucesso!", quiz)
}

type quizFull struct {
	Quizzes         []models.Quiz
	FinishedQuizzes []models.QuizFinished
}

func GetAllQuiz(w http.ResponseWriter, _ *http.Request) {
	quizs, quizsFinisheds, err := service.GetAllQuiz()
	if err != nil {
		log.Printf("Cannot find Get: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar registros, erro interno no sistema.", "")

		return
	} else if len(quizs) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não há questionários cadastrados na base de dados.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Questionários encontrados!", quizFull{Quizzes: quizs, FinishedQuizzes: quizsFinisheds})
}

func PostQuiz(w http.ResponseWriter, r *http.Request) {
	var quiz models.Quiz

	err := json.NewDecoder(r.Body).Decode(&quiz)
	if err != nil {
		log.Printf("Cannot do Post: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Houve algum erro ao tentar obter as informações para cadastro do questionário.", "")

		return
	}

	quiz, err = service.PostQuiz(quiz)
	if err != nil {

		log.Printf("Cannot do Post: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível cadastrar o questionário, houve um erro interno no sistema.", "")

		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Questionário cadastrado com sucesso!", quiz)
}

func PutQuiz(w http.ResponseWriter, r *http.Request) {
	var quiz models.Quiz

	err := json.NewDecoder(r.Body).Decode(&quiz)
	if err != nil {
		log.Printf("Erro ao decodificar os dados do questionário: %v", err)
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Houve um erro ao obter as informações de atualização do questionário.", "")
		return
	}

	log.Println(quiz)

	err = service.PutQuiz(quiz)
	if err != nil {
		log.Printf("Erro ao atualizar o questionário: %v", err)
		statusCode := http.StatusInternalServerError
		message := "Não foi possível atualizar o questionário, ocorreu um erro interno no sistema."
		if strings.Contains(err.Error(), "Questionário já respondido") {
			statusCode = http.StatusBadRequest
		}
		utils.ReturnResponseJSON(w, statusCode, message, "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Informações do questionário atualizadas com sucesso!", "")
}
