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

func GetCboById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Cannot parse ID: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi possível coletar o ID do CBO.", "")
		return
	}

	cbo, err := service.GetCboById(int64(id))
	if err != nil {
		log.Printf("Cannot find Get: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Erro ao buscar por CBO, problema interno no sistema.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "CBO Encontrado com sucesso!", cbo)
}

func GetAllCbo(w http.ResponseWriter, _ *http.Request) {
	cbos, err := service.GetAllCbo()
	if err != nil {
		log.Printf("Cannot find Get: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível procurar pelos CBOs, erro interno no sistema.", "")
		return
	}
	utils.ReturnResponseJSON(w, http.StatusOK, "CBOs Encontrado com sucesso!", cbos)
}

func PostCbo(w http.ResponseWriter, r *http.Request) {
	var cbo models.Cbo

	err := json.NewDecoder(r.Body).Decode(&cbo)
	if err != nil {
		log.Printf("Cannot do Post: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi possível coletar as informações para o registro de CBO.", "")
		return
	}

	cbo, err = service.PostCbo(cbo)
	if err != nil {
		log.Printf("Cannot do Post: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível registrar o CBO, problema interno no sistema.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "CBO registrado com sucesso!", "")
}

func PutCbo(w http.ResponseWriter, r *http.Request) {
	var cbo models.Cbo

	err := json.NewDecoder(r.Body).Decode(&cbo)
	if err != nil {
		log.Printf("Cannot do Put: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi possível coletar as informações para a atualização do CBO.", "")
		return
	}

	cbo, err = service.PutCbo(cbo)
	if err != nil {
		log.Printf("Cannot do Put: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível atualizar o CBO, problema interno no sistema.", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "CBO atualizado com sucesso!", "")
}
