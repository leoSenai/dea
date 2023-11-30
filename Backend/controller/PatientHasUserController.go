package controller

import (
	"api/models"
	"api/service"
	"api/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetPatientHasUserByUserId(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Cannot parse ID: %v", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi possível coletar o id do médico na requisição.", "")
		return
	}

	patientsHasUser, _ := service.GetPatientHasUserByUserId(int64(id))

	utils.ReturnResponseJSON(w, http.StatusOK, "Vinculos com o paciente encontrados com sucesso!", patientsHasUser)
	return
}

func PostPatientHasUser(w http.ResponseWriter, r *http.Request) {
	var patientHasUser models.PatientHasUser

	err := json.NewDecoder(r.Body).Decode(&patientHasUser)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	err = service.PostPatientHasUser(patientHasUser)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível adicionar o paciente, erro interno do servidor.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Paciente adicionado com sucesso!", "")
}
