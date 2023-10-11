package controller

import (
	"api/models"
	"api/models/dtos"
	"api/service"
	"api/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Cannot parse ID: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi especificado o id do usuário procurado.", "")

		return
	}

	user, err := service.GetUserById(int64(id))
	if err != nil {
		log.Printf("Cannot find Get: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar usuário, houve um erro interno no servidor.", "")

		return
	} else if user.IdUser == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não foi possível encontrar o usuário.", "")
		return
	}

	var userDtoSend dtos.UserDTO = dtos.UserDTO{
		IdUser:   user.IdUser,
		Name:     user.Name,
		TypeUser: user.TypeUser,
		IdCbo:    user.IdCbo,
		Active:   user.Active,
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Usuário encontrado com sucesso!", userDtoSend)
}

func GetAllUser(w http.ResponseWriter, _ *http.Request) {
	users, err := service.GetAllUser()
	if err != nil {
		log.Printf("Cannot find Get: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar registros, erro interno no sistema.", "")

		return
	} else if len(users) == 0 {
		utils.ReturnResponseJSON(w, http.StatusNotFound, "Não há usuários cadastrados na base de dados.", "")
		return
	}

	var usersDto []dtos.UserDTO

	for i := 0; i < len(users); i++ {

		usersDto = append(usersDto,
			dtos.UserDTO{
				IdUser:     users[i].IdUser,
				Name:       users[i].Name,
				Active:     users[i].Active,
				IdCbo:      users[i].IdCbo,
				TypeUser:   users[i].TypeUser,
				IdServices: users[i].IdServices,
				Email:      users[i].Email,
				Phone:      users[i].Phone,
			},
		)

	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Usuários encontrados!", usersDto)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Cannot do Post: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Houve algum erro ao tentar obter as informações para cadastro do usuário.", "")

		return
	}

	user, err = service.PostUser(user)
	if err != nil {

		log.Printf("Cannot do Post: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível cadastrar o usuário, houve um erro interno no sistema.", "")

		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Usuário cadastrado com sucesso!", "")
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Cannot do Put: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Houve algum erro ao tentar obter as informações de atualização do usuário.", "")

		return
	}

	user, err = service.PutUser(user)
	if err != nil {

		log.Printf("Cannot do Put: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível atualizar o usuário, houve um erro interno no sistema.", "")

		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Informações do usuário atualizadas com sucesso!", "")
}
