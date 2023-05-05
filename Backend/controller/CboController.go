package controller

import (
	"api/models"
	"api/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetCboById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("Cannot parse ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	cbo, err := service.GetCboById(int64(id))
	if err != nil {
		log.Printf("Cannot find Get: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cbo)
}

func GetAllCbos(w http.ResponseWriter, _ *http.Request) {
	cbos, err := service.GetAllUsers()
	if err != nil {
		log.Printf("Cannot find Get: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cbos)
}

func PostCbo(w http.ResponseWriter, r *http.Request) {
	var cbo models.Cbo

	err := json.NewDecoder(r.Body).Decode(&cbo)
	if err != nil {
		log.Printf("Cannot do Post: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	cbo, err = service.PostCbo(cbo)
	if err != nil {
		log.Printf("Cannot do Post: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cbo)
}

func UpdateCbo(w http.ResponseWriter, r *http.Request) {
	var cbo models.Cbo

	err := json.NewDecoder(r.Body).Decode(&cbo)
	if err != nil {
		log.Printf("Cannot do Put: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	cbo, err = service.PutCbo(cbo)
	if err != nil {
		log.Printf("Cannot do Put: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cbo)
}
