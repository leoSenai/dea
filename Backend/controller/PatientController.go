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

func GetPatientById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Cannot parse ID: %v", err.Error())

		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi possível coletar o id do paciente na requisição.", "")

		return
	}

	patient, err := service.GetPatientById(int64(id))
	if err != nil {
		log.Printf("Cannot find Get: %v", err.Error())
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar o paciente, erro interno no servidor.", "")
		return
	} else if patient.IdPatient == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não foi encontrado pacientes com o ID informado!", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Paciente encontrado com sucesso!", patient)
}

func GetAllPatient(w http.ResponseWriter, _ *http.Request) {
	patients, err := service.GetAllPatient()
	if err != nil {
		log.Printf("Cannot find Get: %v", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar pacientes, erro interno no servidor.", "")

		return
	} else if len(patients) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não foram encontrados registros de Pacientes.", "")
		return
	}
	utils.ReturnResponseJSON(w, http.StatusOK, "Pacientes encontrados com sucesso!", patients)
}

func PostPatient(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient

	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		log.Printf("Cannot do Post: %v", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi possível coletar as informações do paciente necessárias para o cadastro.", "")
		return
	}

	patient, err = service.PostPatient(patient)
	if err != nil {
		log.Printf("Cannot do Post: %v", err.Error())
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível cadastrar o paciente, erro interno no servidor.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Paciente cadastrado com sucesso!", "")
}

func PutPatient(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient

	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		log.Printf("Cannot do Put: %v", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi possível coletar as informações do paciente necessárias para a atualização das informações.", "")
		return
	}

	patient, err = service.PutPatient(patient)
	if err != nil {
		log.Printf("Cannot do Put: %v", err.Error())
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível atualizar as informações do paciente, erro interno no servidor.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Dados do paciente atualizados com sucesso!", "")
}
