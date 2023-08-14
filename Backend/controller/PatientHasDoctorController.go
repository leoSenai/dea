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

func GetPatientHasDoctorByPatientId(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Cannot parse ID: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi especificado o id procurado.", "")

		return
	}

	patientHasDoctor, err := service.GetPatientHasDoctorByPatientId(int64(id))
	if err != nil {
		log.Printf("Cannot find Get: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar os registros, houve um erro interno no servidor.", "")

		return
	} else if len(patientHasDoctor) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não foi possível encontrar os registros.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Registros encontrados com sucesso!", patientHasDoctor)
}

func GetPatientHasDoctorByDoctorId(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Cannot parse ID: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi especificado o id procurado.", "")

		return
	}

	patientHasDoctor, err := service.GetPatientHasDoctorByDoctorId(int64(id))
	if err != nil {
		log.Printf("Cannot find Get: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar os registros, houve um erro interno no servidor.", "")

		return
	} else if len(patientHasDoctor) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não foi possível encontrar os registros.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Registros encontrados com sucesso!", patientHasDoctor)
}

func GetAllPatientHasDoctor(w http.ResponseWriter, _ *http.Request) {
	patientHasDoctors, err := service.GetAllPatientHasDoctor()
	if err != nil {
		log.Printf("Cannot find Get: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar registros, erro interno no sistema.", "")

		return
	} else if len(patientHasDoctors) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não há registros cadastrados na base de dados.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Registros encontrados!", patientHasDoctors)
}

func PostPatientHasDoctor(w http.ResponseWriter, r *http.Request) {
	var patientHasDoctor models.PatientHasDoctor

	err := json.NewDecoder(r.Body).Decode(&patientHasDoctor)
	if err != nil {
		log.Printf("Cannot do Post: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Houve algum erro ao tentar obter as informações para cadastro do registro.", "")

		return
	}

	patientHasDoctor, err = service.PostPatientHasDoctor(patientHasDoctor)
	if err != nil {

		log.Printf("Cannot do Post: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível cadastrar o registro, houve um erro interno no sistema.", "")

		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Registro cadastrado com sucesso!", "")
}
