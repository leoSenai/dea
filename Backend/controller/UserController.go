package controller

import (
	"api/controller/utils"
	"api/models"
	"api/service"
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
		utils.HttpReturnResponseJSON(w, response, 400)

		return
	}

	user, err := service.GetUserById(int64(id))
	if err != nil {
		log.Printf("Cannot find Get: %v", err)

		response := utils.BuildResponseJSON("Não foi possível encontrar usuário, houve um erro interno no servidor.", "")
		utils.HttpReturnResponseJSON(w, response, 500)

		return
	}

	response := utils.BuildResponseJSON("Usuário encontrado com sucesso!", user)
	utils.HttpReturnResponseJSON(w, response, 200)
}

func GetAllUser(w http.ResponseWriter, _ *http.Request) {
	users, err := service.GetAllUser()
	if err != nil {
		log.Printf("Cannot find Get: %v", err)
		http.Error(w, err.Error(), 204)

		response := map[string]interface{}{
			"message": "Não há usuários cadastrados na base de dados.",
			"data":    "",
		}

		utils.HttpReturnResponseJSON(w, response, 400)

		return
	}

	response := map[string]interface{}{
		"message": "Usuários encontrados!",
		"data":    users,
	}

	utils.HttpReturnResponseJSON(w, response, 400)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Cannot do Post: %v", err)
		http.Error(w, err.Error(), 400)

		response := map[string]interface{}{
			"message": "Houve algum erro ao tentar obter as informações para cadastro do usuário.",
			"data":    "",
		}

		utils.HttpReturnResponseJSON(w, response, 400)

		return
	}

	user, err = service.PostUser(user)
	if err != nil {
		log.Printf("Cannot do Post: %v", err)
		http.Error(w, err.Error(), 500)

		response := map[string]interface{}{
			"message": "Não foi possível cadastrar o usuário, houve um erro interno no sistema.",
			"data":    "",
		}

		utils.HttpReturnResponseJSON(w, response, 400)

		return
	}

	response := map[string]interface{}{
		"message": "Usuário cadastrado com sucesso!",
		"data":    "",
	}

	utils.HttpReturnResponseJSON(w, response, 400)
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Cannot do Put: %v", err)
		http.Error(w, err.Error(), 400)

		response := map[string]interface{}{
			"message": "Houve algum erro ao tentar obter as informações de atualização do usuário.",
			"data":    "",
		}

		utils.HttpReturnResponseJSON(w, response, 400)
		return
	}

	user, err = service.PutUser(user)
	if err != nil {
		log.Printf("Cannot do Put: %v", err)
		http.Error(w, err.Error(), 500)

		response := map[string]interface{}{
			"message": "Não foi possível atualizar o usuário, houve um erro interno no sistema.",
			"data":    "",
		}

		utils.HttpReturnResponseJSON(w, response, 400)
		return
	}

	response := map[string]interface{}{
		"message": "Não foi possível atualizar o usuário, houve um erro interno no sistema.",
		"data":    user,
	}

	utils.HttpReturnResponseJSON(w, response, 400)
}
