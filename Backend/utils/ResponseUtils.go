package utils

import (
	"encoding/json"
	"net/http"
)

func ReturnResponseJSON(w http.ResponseWriter, status int, message string, data interface{}) {

	response := map[string]interface{}{
		"message": message,
		"data":    data,
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ReturnResponseFile(w http.ResponseWriter, status int, message string, data http.File) {

	response := map[string]interface{}{
		"message": message,
		"data":    data,
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/pdf")
	json.NewEncoder(w).Encode(response)
}
