package controller

import (
	"api/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func GetPersonById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("Cannot parse ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	person, err := service.GetPersonById(int64(id))
	if err != nil {
		log.Printf("Cannot find Get: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}
