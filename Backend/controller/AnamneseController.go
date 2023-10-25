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

func GetAnamneseById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Cannot parse ID: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi especificado o id da anamnese procurado.", "")

		return
	}

	anamnese, err := service.GetAnamneseById(int64(id))
	if err != nil {
		log.Printf("Cannot find Get: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar a anamnese, houve um erro interno no servidor.", "")

		return
	} else if anamnese.IdAnamnese == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não foi possível encontrar a anamnese.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Anamnese encontrada com sucesso!", anamnese)
}

func GetAllAnamnese(w http.ResponseWriter, _ *http.Request) {
	anamneses, err := service.GetAllAnamnese()
	if err != nil {
		log.Printf("Cannot find Get: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar registros, erro interno no sistema.", "")

		return
	} else if len(anamneses) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não há anamneses cadastradas na base de dados.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Anamneses encontradas!", anamneses)
}

func PostAnamnese(w http.ResponseWriter, r *http.Request) {
	var anamnese models.Anamnese

	err := json.NewDecoder(r.Body).Decode(&anamnese)
	if err != nil {
		log.Printf("Cannot do Post: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Houve algum erro ao tentar obter as informações para cadastro da anamnese.", "")

		return
	}

	anamnese, err = service.PostAnamnese(anamnese)
	if err != nil {

		log.Printf("Cannot do Post: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível cadastrar a anamnese, houve um erro interno no sistema.", "")

		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Anamnese cadastrada com sucesso!", "")
}

func PutAnamnese(w http.ResponseWriter, r *http.Request) {
	var anamnese models.Anamnese

	err := json.NewDecoder(r.Body).Decode(&anamnese)
	if err != nil {
		log.Printf("Cannot do Put: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Houve algum erro ao tentar obter as informações de atualização da anamnese.", "")

		return
	}

	anamnese, err = service.PutAnamnese(anamnese)
	if err != nil {

		log.Printf("Cannot do Put: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível atualizar o anamnese, houve um erro interno no sistema.", "")

		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Informações da anamnese atualizadas com sucesso!", "")
}

func GetAnamneseByUserIdAndPatientId(w http.ResponseWriter, r *http.Request) {
    userIdParam := chi.URLParam(r, "userId")
    patientIdParam := chi.URLParam(r, "patientId")

    userId, err := strconv.Atoi(userIdParam)
    if err != nil {
        log.Printf("Cannot parse userId: %s", err.Error())
        utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi especificado o userId.", "")
        return
    }

    patientId, err := strconv.Atoi(patientIdParam)
    if err != nil {
        log.Printf("Cannot parse patientId: %s", err.Error())
        utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi especificado o patientId.", "")
        return
    }

    anamnese, err := service.GetAnamneseByIdAndPatientId(int64(userId), int64(patientId))
    if err != nil {
        log.Printf("Cannot find Anamnese: %s", err.Error())
        utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar a anamnese, houve um erro interno no servidor.", "")
        return
    } else if anamnese.IdAnamnese == 0 {
        utils.ReturnResponseJSON(w, http.StatusNotFound, "Não foi possível encontrar a anamnese.", "")
        return
    }

    utils.ReturnResponseJSON(w, http.StatusOK, "Anamnese encontrada com sucesso!", anamnese)
}
