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

func GetUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Cannot parse ID: %v", err)

		response := utils.BuildResponseJSON("Não foi especificado o id do usuário procurado.", "")
		utils.ReturnResponseJSON(w, response, 400)

		return
	}

	user, err := service.GetUserById(int64(id))
	if err != nil {
		log.Printf("Cannot find Get: %v", err)

		response := utils.BuildResponseJSON("Não foi possível encontrar usuário, houve um erro interno no servidor.", "")
		utils.ReturnResponseJSON(w, response, 500)

		return
	}

	response := utils.BuildResponseJSON("Usuário encontrado com sucesso!", user)
	utils.ReturnResponseJSON(w, response, 200)
}

func GetAllUser(w http.ResponseWriter, _ *http.Request) {
	users, err := service.GetAllUser()
	if err != nil {
		log.Printf("Cannot find Get: %v", err)
		http.Error(w, err.Error(), 204)

		response := utils.BuildResponseJSON("Não há usuários cadastrados na base de dados.", "")
		utils.ReturnResponseJSON(w, response, 400)

		return
	}

	response := utils.BuildResponseJSON("Usuários encontrados!", users)

	utils.ReturnResponseJSON(w, response, 200)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Cannot do Post: %v", err)
		http.Error(w, err.Error(), 400)

		response := utils.BuildResponseJSON("Houve algum erro ao tentar obter as informações para cadastro do usuário.", "")
		utils.ReturnResponseJSON(w, response, 400)

		return
	}

	user, err = service.PostUser(user)
	if err != nil {
		log.Printf("Cannot do Post: %v", err)
		http.Error(w, err.Error(), 500)

		response := utils.BuildResponseJSON("Não foi possível cadastrar o usuário, houve um erro interno no sistema.", "")
		utils.ReturnResponseJSON(w, response, 400)

		return
	}

	response := utils.BuildResponseJSON("Usuário cadastrado com sucesso!", "")
	utils.ReturnResponseJSON(w, response, 200)
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Cannot do Put: %v", err)
		http.Error(w, err.Error(), 400)

		response := utils.BuildResponseJSON("Houve algum erro ao tentar obter as informações de atualização do usuário.", "")
		utils.ReturnResponseJSON(w, response, 400)

		return
	}

	user, err = service.PutUser(user)
	if err != nil {
		log.Printf("Cannot do Put: %v", err)
		http.Error(w, err.Error(), 500)

		response := utils.BuildResponseJSON("Não foi possível atualizar o usuário, houve um erro interno no sistema.", "")
		utils.ReturnResponseJSON(w, response, 400)

		return
	}

	response := map[string]interface{}{
		"message": "Informações do usuário atualizadas com sucesso!",
		"data":    "",
	}

	utils.ReturnResponseJSON(w, response, 200)
}
