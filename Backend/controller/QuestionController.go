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

func GetQuestionById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Cannot parse ID: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi especificado o id da questão procurado.", "")

		return
	}

	question, err := service.GetQuestionById(int64(id))
	if err != nil {
		log.Printf("Cannot find Get: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar a questão, houve um erro interno no servidor.", "")

		return
	} else if question.IdQuestion == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não foi possível encontrar a questão.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Questão encontrada com sucesso!", question)
}

func GetAllQuestion(w http.ResponseWriter, _ *http.Request) {
	questions, err := service.GetAllQuestion()
	if err != nil {
		log.Printf("Cannot find Get: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar registros, erro interno no sistema.", "")

		return
	} else if len(questions) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não há questões cadastradas na base de dados.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Questões encontradas!", questions)
}

func PostQuestion(w http.ResponseWriter, r *http.Request) {
	var question models.Question

	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		log.Printf("Cannot do Post: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Houve algum erro ao tentar obter as informações para cadastro da questão.", "")

		return
	}

	question, err = service.PostQuestion(question)
	if err != nil {

		log.Printf("Cannot do Post: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível cadastrar a questão, houve um erro interno no sistema.", "")

		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Questão cadastrado com sucesso!", "")
}

func PutQuestion(w http.ResponseWriter, r *http.Request) {
	var question models.Question

	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		log.Printf("Cannot do Put: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Houve algum erro ao tentar obter as informações de atualização da questão.", "")

		return
	}

	question, err = service.PutQuestion(question)
	if err != nil {

		log.Printf("Cannot do Put: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível atualizar a questão, houve um erro interno no sistema.", "")

		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Informações da questão atualizadas com sucesso!", "")
}
