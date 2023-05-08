package controller

import (
	"api/models"
	"api/service"
	"api/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func GetPersonById(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if id <= 0 {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "ID inválido", "")
		return
	}

	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	person, err := service.GetPersonById(int64(id))
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, err.Error(), "")
		return
	}

	if person.IdPerson == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Pessoa não encontrada", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Pessoa encontrada com sucesso", person)
}

func PostPerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	err = service.PostPerson(person)
	if err != nil {
		var status int
		var message string

		if strings.Contains(err.Error(), "já cadastrado!") {
			status = http.StatusBadRequest
			message = err.Error()
		} else {
			status = http.StatusInternalServerError
			message = http.StatusText(http.StatusInternalServerError)
		}

		utils.ReturnResponseJSON(w, status, message, "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Pessoa cadastrada com sucesso!", "")
}

func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	persons, err := service.GetAllPerson()
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	if len(persons) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não há pessoas cadastradas ainda.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Busca realizada com sucesso!", persons)
}

func PutPerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	err = service.PutPerson(person)
	if err != nil {
		if strings.Contains(err.Error(), "não está cadastrada") {
			utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
			return
		} else if strings.Contains(err.Error(), "Nenhum dado foi atualizado") {
			utils.ReturnResponseJSON(w, http.StatusCreated, err.Error(), "")
			return
		}

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, err.Error(), "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Pessoa atualizada com sucesso!", "")
}
