package controller

import (
	"api/models"
	"api/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func HttpReturnResponseJSON(w http.ResponseWriter, response map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Cannot parse ID: %v", err)
		http.Error(w, err.Error(), 400)

		response := map[string]interface{}{
			"message": "Não foi especificado o id do usuário procurado.",
			"data":    "",
		}

		HttpReturnResponseJSON(w, response)

		return
	}

	user, err := service.GetUserById(int64(id))
	if err != nil {
		log.Printf("Cannot find Get: %v", err)
		http.Error(w, err.Error(), 500)

		response := map[string]interface{}{
			"message": "Não foi possível encontrar usuário, houve um erro interno no servidor.",
			"data":    "",
		}

		HttpReturnResponseJSON(w, response)

		return
	}

	response := map[string]interface{}{
		"message": "Usuário encontrado com sucesso!",
		"data":    user,
	}

	HttpReturnResponseJSON(w, response)
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

		HttpReturnResponseJSON(w, response)

		return
	}

	response := map[string]interface{}{
		"message": "Usuários encontrados!",
		"data":    users,
	}

	HttpReturnResponseJSON(w, response)
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

		HttpReturnResponseJSON(w, response)

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

		HttpReturnResponseJSON(w, response)

		return
	}

	response := map[string]interface{}{
		"message": "Usuário cadastrado com sucesso!",
		"data":    "",
	}

	HttpReturnResponseJSON(w, response)
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

		HttpReturnResponseJSON(w, response)
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

		HttpReturnResponseJSON(w, response)
		return
	}

	response := map[string]interface{}{
		"message": "Não foi possível atualizar o usuário, houve um erro interno no sistema.",
		"data":    user,
	}

	HttpReturnResponseJSON(w, response)
}
