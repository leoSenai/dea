package controller

import (
	"api/models"
	"api/service"
	"api/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func PostProximity(w http.ResponseWriter, r *http.Request) {
	var proximity models.Proximity

	err := json.NewDecoder(r.Body).Decode(&proximity)

	if err != nil {
		log.Println("err", err)
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}

	err = service.PostProximity(proximity)
	if err != nil {
		var status int
		var message string

		if strings.Contains(err.Error(), "Pessoa não cadastrada") ||
			strings.Contains(err.Error(), "Paciente não cadastrado") ||
			strings.Contains(err.Error(), "já cadastrado!") {
			status = http.StatusBadRequest
			message = err.Error()
		} else {
			status = http.StatusInternalServerError
			message = err.Error()
		}

		utils.ReturnResponseJSON(w, status, message, "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Proximidade cadastrado com sucesso!", "")
}

func GetProximityAllByIdPerson(w http.ResponseWriter, r *http.Request) {
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

	proximitys, err := service.GetProximityAllByIdPerson(int64(id))
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, err.Error(), "")
		return
	}

	if len(proximitys) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não há proximidades cadastradas a esta pessoa", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Proximidades encontrada com sucesso!", proximitys)
}

func GetProximityAllByIdPatient(w http.ResponseWriter, r *http.Request) {
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

	proximitys, err := service.GetProximityAllByIdPatient(int64(id))
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, err.Error(), "")
		return
	}

	if len(proximitys) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não há proximidades cadastradas a este paciente", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Proximidades encontrada com sucesso!", proximitys)
}
