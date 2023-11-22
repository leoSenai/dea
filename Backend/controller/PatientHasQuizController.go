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

func GetPatientQuizByQuizID(w http.ResponseWriter, r *http.Request) {
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

	var patientHasQuiz []models.PatientHasQuiz

	patientHasQuiz, err = service.GetPatientQuizByQuizID(int64(id))

	if len(patientHasQuiz) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNoContent, "Não existem pacientes para este questionario", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Busca realizada com sucesso!", patientHasQuiz)
}

func GetPatientQuizByQuizPatientID(w http.ResponseWriter, r *http.Request) {
	idQuizParam := chi.URLParam(r, "idquiz")
	idPatientParam := chi.URLParam(r, "idpatient")

	idQuiz, err := strconv.Atoi(idQuizParam)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "Ops! Parâmetros incorretos.")
		return
	}
	idPatient, err := strconv.Atoi(idPatientParam)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "Ops! Parâmetros incorretos.")
		return
	}

	if idQuiz <= 0 || idPatient <= 0 {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "ID inválido", "")
		return
	}

	var patientHasQuiz []models.PatientHasQuiz

	_, patientHasQuiz, err = service.GetPatientQuizByQuizPatientID(int64(idQuiz), int64(idPatient))
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, err.Error(), "Ops! Erro interno do servidor.")
		return
	}

	if len(patientHasQuiz) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não existem relacionamentos com este ID de questionário e paciente", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Busca realizada com sucesso!", patientHasQuiz)

}

func GetPatientQuizByPatientID(w http.ResponseWriter, r *http.Request) {
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

	var patientHasQuiz []models.PatientHasQuiz

	patientHasQuiz, err = service.GetPatientQuizByPatientID(int64(id))

	if len(patientHasQuiz) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não existem relacionamento com este ID de Paciente", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Busca realizada com sucesso!", patientHasQuiz)
}

func PostPatientQuiz(w http.ResponseWriter, r *http.Request) {
	var patientHasQuiz models.PatientHasQuiz

	err := json.NewDecoder(r.Body).Decode(&patientHasQuiz)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	err = service.PostPatientQuiz(patientHasQuiz)
	if err != nil {
		var status int
		var message string

		if strings.Contains(err.Error(), "Já esta relacionado no banco de dados") {
			status = http.StatusBadRequest
			message = err.Error()
		} else {
			status = http.StatusInternalServerError
			message = http.StatusText(http.StatusInternalServerError)
		}

		utils.ReturnResponseJSON(w, status, message, "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Relacionamento cadastrado com sucesso", "")
}

func PutPatientQuiz(w http.ResponseWriter, r *http.Request) {
	var patientHasQuiz models.PatientHasQuiz

	err := json.NewDecoder(r.Body).Decode(&patientHasQuiz)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	err = service.PutPatientQuiz(patientHasQuiz)
	if err != nil {
		var status int
		var message string
		if strings.Contains(err.Error(), "Relacionamento não cadastrado no banco de dados") {
			status = http.StatusBadRequest
			message = err.Error()
		} else {
			status = http.StatusInternalServerError
			message = http.StatusText(http.StatusInternalServerError)
		}

		utils.ReturnResponseJSON(w, status, message, "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Relacionamento atualizado com sucesso!", "")
}
