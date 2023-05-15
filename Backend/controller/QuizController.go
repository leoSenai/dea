package controller

import (
	"api/models"
	"api/service"
	"api/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

func GetAllQuiz(w http.ResponseWriter, _ *http.Request) {
	quizs, err := service.GetAllQuiz()
	if err != nil {
		log.Printf("Cannot find Get: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar registros, erro interno no sistema.", "")

		return
	} else if len(quizs) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNoContent, "Não há questionários cadastrados na base de dados.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Questionários encontrados!", quizs)
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

	utils.ReturnResponseJSON(w, http.StatusOK, "Questionário cadastrado com sucesso!", "")
}

func PutQuiz(w http.ResponseWriter, r *http.Request) {
	var quiz models.Quiz

	err := json.NewDecoder(r.Body).Decode(&quiz)
	if err != nil {
		log.Printf("Cannot do Put: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Houve algum erro ao tentar obter as informações de atualização do questionário.", "")

		return
	}

	quiz, err = service.PutQuiz(quiz)
	if err != nil {

		log.Printf("Cannot do Put: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível atualizar o questionário, houve um erro interno no sistema.", "")

		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Informações do questionário atualizadas com sucesso!", "")
}
