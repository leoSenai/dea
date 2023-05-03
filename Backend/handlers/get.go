package handlers

import (
	"api/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	pessoa, err := models.Get(int64(id))
	if err != nil {
		log.Printf("Erro ao buscar Get: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pessoa)
}

func GetTeste(w http.ResponseWriter, r *http.Request) {
	models.Hello()
}
