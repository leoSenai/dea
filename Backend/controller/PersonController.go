package controller

import (
	"api/models"
	"api/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func GetPersonById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	person, err := service.GetPersonById(int64(id))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func PostPerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = service.InsertPerson(person)
	if err != nil {
		if strings.Contains(err.Error(), "já cadastrado!") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"error":   false,
		"message": "Pessoa cadastrado com sucesso!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	persons, err := service.GetAllPersons()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&persons)
}

func PutPerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = service.UpdatePerson(person)
	if err != nil {
		if strings.Contains(err.Error(), "não está cadastrada") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"error":   false,
		"message": "Pessoa autualizado com sucesso!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
