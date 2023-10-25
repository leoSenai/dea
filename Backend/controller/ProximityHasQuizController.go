package controller

import (
	"api/models"
	"api/models/dtos"
	"api/service"
	"api/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func GetProximityQuizByQuizPatientPersonIDs(w http.ResponseWriter, r *http.Request) {
	var proximityHasQuizDto dtos.ProximityHasQuizDto

	err := json.NewDecoder(r.Body).Decode(&proximityHasQuizDto)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	var proximityHasQuiz models.ProximityHasQuiz
	proximityHasQuiz, err = service.GetProximityQuizByQuizPatientPersonIDs(proximityHasQuizDto)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, err.Error(), "")
		return
	}

	if proximityHasQuiz.IdQuiz == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Relacionamento em proximidade e questionario não encontrado", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Relacionamento em proximidade e questionario encontrada com sucesso", proximityHasQuiz)
}

func GetProximityQuizByQuizID(w http.ResponseWriter, r *http.Request) {
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

	var proximityHasQuiz []models.ProximityHasQuiz

	proximityHasQuiz, err = service.GetProximityQuizByQuizID(int64(id))

	if len(proximityHasQuiz) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não existem relacionamento para este questionario", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Busca realizada com sucesso!", proximityHasQuiz)
}

func GetProximityQuizByPatientID(w http.ResponseWriter, r *http.Request) {
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

	var proximityHasQuiz []models.ProximityHasQuiz

	proximityHasQuiz, err = service.GetProximityQuizByPatientID(int64(id))

	if len(proximityHasQuiz) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não existem relacionamento com este ID de Paciente", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Busca realizada com sucesso!", proximityHasQuiz)
}

func GetProximityQuizByPersonID(w http.ResponseWriter, r *http.Request) {
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

	var proximityHasQuiz []models.ProximityHasQuiz

	proximityHasQuiz, err = service.GetProximityQuizByPersonID(int64(id))

	if len(proximityHasQuiz) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não existem relacionamento com este ID de Pessoa", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Busca realizada com sucesso!", proximityHasQuiz)
}

func PostProximityQuiz(w http.ResponseWriter, r *http.Request) {
	var proximityHasQuiz models.ProximityHasQuiz

	err := json.NewDecoder(r.Body).Decode(&proximityHasQuiz)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	err = service.PostProximityQuiz(proximityHasQuiz)
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

func PutProximityQuiz(w http.ResponseWriter, r *http.Request) {
	var proximityHasQuiz models.ProximityHasQuiz

	err := json.NewDecoder(r.Body).Decode(&proximityHasQuiz)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	err = service.PutProximityQuiz(proximityHasQuiz)
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
