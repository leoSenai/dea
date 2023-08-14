package controller

import (
	"api/models"
	"api/service"
	"api/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func GetAskingById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if id <= 0 {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "ID inválido", "")
		return
	}

	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	asking, err := service.GetAskingById(int64(id))
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, err.Error(), "")
		return
	}

	if asking.IdAsking == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Pergunta não encontrada", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Pergunta encontrada com sucesso", asking)
}

func PostAsking(w http.ResponseWriter, r *http.Request) {
	var asking models.Asking

	err := json.NewDecoder(r.Body).Decode(&asking)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	err = service.PostAsking(asking)
	if err != nil {
		var status int
		var message string

		if strings.Contains(err.Error(), "já cadastrado!") {
			status = http.StatusBadRequest
			message = err.Error()
		} else {
			status = http.StatusInternalServerError
			message = http.StatusText(http.StatusInternalServerError)
		}

		utils.ReturnResponseJSON(w, status, message, "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Pergunta cadastrado com sucesso!", "")
}

func GetAllAsking(w http.ResponseWriter, _ *http.Request) {
	askings, err := service.GetAllAsking()
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	if len(askings) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não há médicos cadastrados ainda.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Busca realizada com sucesso!", askings)
}

func PutAsking(w http.ResponseWriter, r *http.Request) {
	var asking models.Asking

	err := json.NewDecoder(r.Body).Decode(&asking)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	err = service.PutAsking(asking)
	if err != nil {
		if strings.Contains(err.Error(), "não está cadastrado") {
			utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
			return
		} else if strings.Contains(err.Error(), "Nenhum dado foi atualizado") {
			utils.ReturnResponseJSON(w, http.StatusCreated, err.Error(), "")
			return
		}

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, err.Error(), "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Pergunta atualizada com sucesso!", "")
}
