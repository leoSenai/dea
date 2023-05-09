package utils

import (
	"encoding/json"
	"net/http"
)

func HttpReturnResponseJSON(w http.ResponseWriter, response map[string]interface{}, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func BuildResponseJSON(message string, data interface{}) (response map[string]interface{}) {
	response = map[string]interface{}{
		"message": message,
		"data":    data,
	}
	return response
}
