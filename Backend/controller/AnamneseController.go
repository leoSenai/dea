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

func GetAnamneseByIdUserPatient(w http.ResponseWriter, r *http.Request) {
	idUserParam := chi.URLParam(r, "idUser")
	idPatientParam := chi.URLParam(r, "idPatient")
	idUser, err := strconv.Atoi(idUserParam)
	if err != nil {
		log.Printf("Cannot parse ID User: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi especificado o id do médico procurado.", "")

		return
	}
	idPatient, err := strconv.Atoi(idPatientParam)
	if err != nil {
		log.Printf("Cannot parse ID Patient: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi especificado o id do paciente procurado.", "")

		return
	}

	anamnese, err := service.GetAnamneseByIdUserPatient(int64(idUser), int64(idPatient))
	if err != nil {
		log.Printf("Cannot find Get: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível encontrar a anamnese, houve um erro interno no servidor.", "")

		return
	} else if anamnese.IdAnamnese == 0 {
		utils.ReturnResponseJSON(w, http.StatusNoContent, "Não foi possível encontrar a anamnese.", "")
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

	anamneseReturned, _ := service.GetAnamneseByIdUserPatient(int64(anamnese.IdUser), int64(anamnese.IdPatient))
	if anamneseReturned.IdAnamnese != 0 {
		anamneseReturned.Notes = anamnese.Notes
		anamneseReturned.Indicative = anamnese.Indicative
		anamnese, err = service.PutAnamnese(anamneseReturned)
	} else {
		anamnese, err = service.PostAnamnese(anamnese)
	}

	if err != nil {

		log.Printf("Cannot do Put: %s", err.Error())

		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Não foi possível atualizar o anamnese, houve um erro interno no sistema.", "")

		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Informações da anamnese atualizadas com sucesso!", "")
}

func GetSumResults(w http.ResponseWriter, r *http.Request) {
	var sumresults []models.SumResult
	idPatientParam := chi.URLParam(r, "idpatient")

	idPatient, err := strconv.Atoi(idPatientParam)
	if err != nil {
		log.Printf("Cannot parse ID Patient: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi especificado o id do paciente.", "")

		return
	}
	sumresults, _ = service.GetSumResults(idPatient)

	utils.ReturnResponseJSON(w, http.StatusOK, "Resultados encontrados!", sumresults)
	return
}

func GetReport(w http.ResponseWriter, r *http.Request) {
	var report http.File

	idUserParam := chi.URLParam(r, "iduser")
	idPatientParam := chi.URLParam(r, "idpatient")
	grau := chi.URLParam(r, "grau")

	id_user, err := strconv.Atoi(idUserParam)
	if err != nil {
		log.Printf("Cannot parse ID User: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi especificado o id do médico no corpo de requisição.", "")
		return
	}
	id_patient, err := strconv.Atoi(idPatientParam)
	if err != nil {
		log.Printf("Cannot parse ID User: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi especificado o id do paciente no corpo de requisição.", "")
		return
	}

	grau_int, err := strconv.Atoi(grau)
	if err != nil {
		log.Printf("Cannot parse ID User: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi especificado o id do paciente no corpo de requisição.", "")
		return
	}

	var anamnese_model models.Anamnese
	anamnese_model.IdUser = id_user
	anamnese_model.IdPatient = id_patient

	anamnese_model, err = service.GetAnamneseByIdUserPatient(int64(anamnese_model.IdUser), int64(id_patient))
	if err != nil {
		log.Printf("Cannot parse ID User: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi possível buscar informações da anamnese.", "")

		return
	}

	report, err = service.GetReport(anamnese_model, grau_int)
	if err != nil {
		log.Printf("Cannot parse ID User: %s", err.Error())
		utils.ReturnResponseJSON(w, http.StatusBadRequest, "Não foi possível gerar o laudo da anamnese.", "")

		return
	}

	utils.ReturnResponseFile(w, http.StatusOK, "Laudo gerado! Aguarde o download do laudo começar...", report)
	return
}
