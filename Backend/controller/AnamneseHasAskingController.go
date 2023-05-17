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

func GetAnamneseHasAskingByAnamneseId(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Cannot parse ID: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi especificado o id procurado.", "")

		return
	}

	anamneseHasAsking, err := service.GetAnamneseHasAskingByAnamneseId(int64(id))
	if err != nil {
		log.Printf("Cannot find Get: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar o registro, houve um erro interno no servidor.", "")

		return
	} else if anamneseHasAsking.IdAnamnese == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não foi possível encontrar o registro.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Registro encontrado com sucesso!", anamneseHasAsking)
}

func GetAnamneseHasAskingByAskingId(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Cannot parse ID: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi especificado o id procurado.", "")

		return
	}

	anamneseHasAsking, err := service.GetAnamneseHasAskingByAskingId(int64(id))
	if err != nil {
		log.Printf("Cannot find Get: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar o registro, houve um erro interno no servidor.", "")

		return
	} else if anamneseHasAsking.IdAnamnese == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não foi possível encontrar o registro.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Registro encontrado com sucesso!", anamneseHasAsking)
}

func GetAllAnamneseHasAsking(w http.ResponseWriter, _ *http.Request) {
	anamneseHasAskings, err := service.GetAllAnamneseHasAsking()
	if err != nil {
		log.Printf("Cannot find Get: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar registros, erro interno no sistema.", "")

		return
	} else if len(anamneseHasAskings) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não há registros cadastrados na base de dados.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Registros encontrados!", anamneseHasAskings)
}

func PostAnamneseHasAsking(w http.ResponseWriter, r *http.Request) {
	var anamneseHasAsking models.AnamneseHasAsking

	err := json.NewDecoder(r.Body).Decode(&anamneseHasAsking)
	if err != nil {
		log.Printf("Cannot do Post: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Houve algum erro ao tentar obter as informações para cadastro do registro.", "")

		return
	}

	anamneseHasAsking, err = service.PostAnamneseHasAsking(anamneseHasAsking)
	if err != nil {

		log.Printf("Cannot do Post: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível cadastrar o registro, houve um erro interno no sistema.", "")

		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Registro cadastrado com sucesso!", "")
}

func DeleteAnamneseHasAsking(w http.ResponseWriter, r *http.Request) {
	var anamneseHasAsking models.AnamneseHasAsking

	err := json.NewDecoder(r.Body).Decode(&anamneseHasAsking)
	if err != nil {
		log.Printf("Cannot do Put: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Houve algum erro ao tentar obter as informações de remoção do registro.", "")

		return
	}

	anamneseHasAsking, err = service.DeleteAnamneseHasAsking(anamneseHasAsking)
	if err != nil {

		log.Printf("Cannot do Put: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível atualizar o usuário, houve um erro interno no sistema.", "")

		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Informações do usuário atualizadas com sucesso!", "")
}
