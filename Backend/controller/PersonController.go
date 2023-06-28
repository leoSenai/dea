package controller

import (
	"api/models"
	"api/models/dtos"
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

	var personDto dtos.PersonDTO = dtos.PersonDTO{
		IdPerson:  person.IdPerson,
		Name:      person.Name,
		BornDate:  person.BornDate,
		DocNumber: person.DocNumber,
		DocType:   person.DocType,
		Email:     person.Email,
		Password:  "",
		Salt:      "",
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Pessoa encontrada com sucesso", personDto)
}

func PostPerson(w http.ResponseWriter, r *http.Request) {
	var personDto dtos.PersonDTO

	err := json.NewDecoder(r.Body).Decode(&personDto)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	err = service.PostPerson(personDto)
	if err != nil {
		var status int
		var message string

		if strings.Contains(err.Error(), "já cadastrado!") {
			status = http.StatusBadRequest
			message = err.Error()
		} else if strings.Contains(err.Error(), "Não foi possivel") {
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

func GetAllPerson(w http.ResponseWriter, _ *http.Request) {
	persons, err := service.GetAllPerson()
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	if len(persons) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não há pessoas cadastradas ainda.", "")
		return
	}

	var personsDto []dtos.PersonDTO

	for i := 0; i < len(persons); i++ {
		personsDto = append(personsDto, dtos.PersonDTO{
			IdPerson:  persons[i].IdPerson,
			Name:      persons[i].Name,
			BornDate:  persons[i].BornDate,
			DocNumber: persons[i].DocNumber,
			DocType:   persons[i].DocType,
			Email:     persons[i].Email,
			Password:  "",
			Salt:      "",
		})
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Busca realizada com sucesso!", personsDto)
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

func GetPersonByDocNumber(w http.ResponseWriter, r *http.Request) {
	docNumber := chi.URLParam(r, "docNumber")

	person, err := service.GetPersonByDocNumber(docNumber)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, err.Error(), "")
		return
	}

	if person.IdPerson == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Pessoa não encontrada", "")
		return
	}

	var personDto dtos.PersonDTO = dtos.PersonDTO{
		IdPerson:  person.IdPerson,
		Name:      person.Name,
		BornDate:  person.BornDate,
		DocNumber: person.DocNumber,
		DocType:   person.DocType,
		Email:     person.Email,
		Password:  "",
		Salt:      "",
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Pessoa encontrada com sucesso", personDto)
}
