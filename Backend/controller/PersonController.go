package controller

import (
	"api/models"
	"api/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func GetPersonById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	person, err := service.GetPersonById(int64(id))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Pessoa autualizado com sucesso!",
		"data":    json.NewEncoder(w).Encode(person),
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func PostPerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.PostPerson(person)
	if err != nil {
		var status int
		var message string
		if strings.Contains(err.Error(), "já cadastrado!") {
			status = 400
			message = err.Error()
		} else {
			status = 500
			message = http.StatusText(http.StatusInternalServerError)
		}

		response := map[string]interface{}{
			"message": message,
			"data":    "",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := map[string]interface{}{
		"message": "Pessoa cadastrado com sucesso!",
		"data":    "",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetAllPerson(w http.ResponseWriter, _ *http.Request) {
	persons, err := service.GetAllPerson()
	var status int
	var message string
	var data []models.Person
	if err != nil {
		status = 400
		message = err.Error()
	} else if len(persons) == 0 {
		status = 201
		message = "Não há pessoas cadastradas ainda."
	} else {
		status = 200
		message = "Busca realizada com sucesso!"
		data = persons
	}
	response := map[string]interface{}{
		"message": message,
		"data":    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func PutPerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = service.PutPerson(person)
	if err != nil {
		if strings.Contains(err.Error(), "não está cadastrada") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "Pessoa autualizado com sucesso!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
