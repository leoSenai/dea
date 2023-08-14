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

func GetDoctorById(w http.ResponseWriter, r *http.Request) {
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

	doctor, err := service.GetDoctorById(int64(id))
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, err.Error(), "")
		return
	}

	if doctor.IdDoctor == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Médico não encontrada", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Médico encontrada com sucesso", doctor)
}

func PostDoctor(w http.ResponseWriter, r *http.Request) {
	var doctor models.Doctor

	err := json.NewDecoder(r.Body).Decode(&doctor)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	err = service.PostDoctor(doctor)
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

	utils.ReturnResponseJSON(w, http.StatusOK, "Médico cadastrado com sucesso!", "")
}

func GetAllDoctor(w http.ResponseWriter, _ *http.Request) {
	doctors, err := service.GetAllDoctor()
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	if len(doctors) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não há médicos cadastrados ainda.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Busca realizada com sucesso!", doctors)
}

func PutDoctor(w http.ResponseWriter, r *http.Request) {
	var doctor models.Doctor

	err := json.NewDecoder(r.Body).Decode(&doctor)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	err = service.PutDoctor(doctor)
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

	utils.ReturnResponseJSON(w, http.StatusOK, "Médico atualizada com sucesso!", "")
}
