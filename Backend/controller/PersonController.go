package controller

import (
	"api/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	pessoa, err := service.GetPersonById(int64(id))
	if err != nil {
		log.Printf("Erro ao buscar Get: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pessoa)
}
