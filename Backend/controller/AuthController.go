package controller

import (
	"api/models"
	"api/service"
	"api/utils"
	"encoding/json"
	"net/http"
	"strings"
)

func PostLogin(w http.ResponseWriter, r *http.Request) {
	var login models.Login

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}
	var token string
	token, err = service.PostLogin(login)
	if err != nil {
		var status int
		var message string

		if strings.Contains(err.Error(), "credenciais sao invalidas") {
			status = http.StatusBadRequest
			message = err.Error()
		} else {
			status = http.StatusInternalServerError
			message = http.StatusText(http.StatusInternalServerError)
		}

		utils.ReturnResponseJSON(w, status, message, "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Credenciais válidas!", token)
}
