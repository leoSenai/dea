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

func GetServicesById(w http.ResponseWriter, r *http.Request) {
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

	services, err := service.GetServicesById(int64(id))
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, err.Error(), "")
		return
	}

	if services.IdServices == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Serviço não encontrada", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Serviço encontrada com sucesso", services)
}

func PostServices(w http.ResponseWriter, r *http.Request) {
	var services models.Services

	err := json.NewDecoder(r.Body).Decode(&services)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	err = service.PostServices(services)
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

	utils.ReturnResponseJSON(w, http.StatusOK, "Serviço cadastrado com sucesso!", "")
}

func GetAllServices(w http.ResponseWriter, _ *http.Request) {
	servicess, err := service.GetAllServices()
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	if len(servicess) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não há médicos cadastrados ainda.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Busca realizada com sucesso!", servicess)
}

func PutServices(w http.ResponseWriter, r *http.Request) {
	var services models.Services

	err := json.NewDecoder(r.Body).Decode(&services)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	err = service.PutServices(services)
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

	utils.ReturnResponseJSON(w, http.StatusOK, "Serviço atualizada com sucesso!", "")
}
