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
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Cannot parse ID: %v", err)

		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi possível coletar o id do paciente na requisição.", "")

		return
	}

	patient, err := service.GetPatientById(int64(id))
	if err != nil {
		log.Printf("Cannot find Get: %v", err)
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar o paciente, erro interno no servidor.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Paciente encontrado com sucesso!", patient)
}

func GetAllPatient(w http.ResponseWriter, _ *http.Request) {
	patients, err := service.GetAllPatient()
	if err != nil {
		log.Printf("Cannot find Get: %v", err)

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar pacientes, erro interno no servidor.", "")

		return
	}
	utils.ReturnResponseJSON(w, http.StatusOK, "Pacientes encontrados com sucesso!", patients)
}

func PostPatient(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient

	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		log.Printf("Cannot do Post: %v", err)
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi possível coletar as informações do paciente necessárias para o cadastro.", "")
		return
	}

	patient, err = service.PostPatient(patient)
	if err != nil {
		log.Printf("Cannot do Post: %v", err)
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível cadastrar o paciente, erro interno no servidor.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Paciente cadastrado com sucesso!", patient)
}

func PutPatient(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient

	err := json.NewDecoder(r.Body).Decode(&patient)
	if err != nil {
		log.Printf("Cannot do Put: %v", err)
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi possível coletar as informações do paciente necessárias para a atualização das informações.", "")
		return
	}

	patient, err = service.PutPatient(patient)
	if err != nil {
		log.Printf("Cannot do Put: %v", err)
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível atualizar as informações do paciente, erro interno no servidor.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Dados do paciente atualizados com sucesso!", patient)
}
