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

	json.NewEncoder(w).Encode(response)

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
