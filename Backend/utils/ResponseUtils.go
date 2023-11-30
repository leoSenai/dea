package utils

import (
	"encoding/base64"
	"encoding/json"
	"log"
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

func ReturnResponseFile(w http.ResponseWriter, status int, message string, data []byte) {
	log.Println(data)
	response := map[string]interface{}{
		"message": message,
		"data":    base64.StdEncoding.EncodeToString(data),
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/pdf")
	json.NewEncoder(w).Encode(response)

}
