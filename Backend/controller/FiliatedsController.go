package controller

import (
	"api/models"
	"api/service"
	"api/utils"
	"encoding/json"
	"net/http"
)

type FiliatedsReceived struct {
	Filiateds        []models.Filiated
	FiliatedsRemoved []models.Filiated
}

func VerifyIfPatientHasQuizExist(patientHasQuizModel models.PatientHasQuiz) (alreadyExist bool, patientHasQuizzes []models.PatientHasQuiz) {

	alreadyExist, patientHasQuizzes, _ = service.GetPatientQuizByQuizPatientID(patientHasQuizModel.IdQuiz, patientHasQuizModel.IdPatient)

	return alreadyExist, patientHasQuizzes
}

func VerifyIfProximityHasQuizExist(proximityHasQuizModel models.ProximityHasQuiz) (alreadyExist bool, proximityHasQuizzes []models.ProximityHasQuiz) {

	alreadyExist, proximityHasQuizzes, _ = service.GetProximityQuizByQuizPersonID(proximityHasQuizModel.IdQuiz, proximityHasQuizModel.ProximityIdPerson)

	return alreadyExist, proximityHasQuizzes
}

func ConvertFiliatedsToPatientOrPersonQuizzes(filiateds []models.Filiated) (patientHasQuiz []models.PatientHasQuiz, proximityHasQuiz []models.ProximityHasQuiz, err error) {

	var patientHasQuizModel models.PatientHasQuiz
	var proximityHasQuizModel models.ProximityHasQuiz

	for _, filiated := range filiateds {
		if filiated.Type == "Patient" { //Patient

			patientHasQuizModel.IdQuiz = filiated.IdQuiz

			patientDb, _ := service.GetPatientByName(filiated.Name)

			patientHasQuizModel.IdPatient = int64(patientDb.IdPatient)
			patientHasQuizModel.Finished = 0

			patientHasQuiz = append(patientHasQuiz, patientHasQuizModel)

		} else { //Person

			proximityHasQuizModel.IdQuiz = filiated.IdQuiz

			patientDb, _ := service.GetPersonByName(filiated.Name)

			proximityHasQuizModel.ProximityIdPerson = patientDb.IdPerson
			proximityHasQuizModel.Finished = 0

			proximityHasQuiz = append(proximityHasQuiz, proximityHasQuizModel)

		}
	}

	return
}

func UpdatePatientHasQuizzes(patientsHasQuiz []models.PatientHasQuiz, w http.ResponseWriter) {
	for _, patientHasQuiz := range patientsHasQuiz {
		alreadyExist, _ := VerifyIfPatientHasQuizExist(patientHasQuiz)

		if !alreadyExist {
			err := service.PostPatientQuiz(patientHasQuiz)
			if err != nil {
				utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Erro interno! Não foi possível cadastrar o paciente ao questionário.", "")
				return
			}
		}
	}
}

func UpdateProximityHasQuizzes(proximitysHasQuiz []models.ProximityHasQuiz, w http.ResponseWriter) {
	for _, proximityHasQuiz := range proximitysHasQuiz {
		alreadyExist, _ := VerifyIfProximityHasQuizExist(proximityHasQuiz)

		if !alreadyExist {
			err := service.PostProximityQuiz(proximityHasQuiz)
			if err != nil {
				utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Erro interno! Não foi possível cadastrar o paciente ao questionário.", "")
				return
			}
		}
	}
}

// remove funcs
func RemovePatientHasQuizzes(patientsHasQuizRemoved []models.PatientHasQuiz, w http.ResponseWriter) {
	for _, patientHasQuizRemoved := range patientsHasQuizRemoved {
		alreadyExist, patientHasQuizRemovedFull := VerifyIfPatientHasQuizExist(patientHasQuizRemoved)

		if alreadyExist {
			patientHasQuizRemoved = patientHasQuizRemovedFull[0]
			err := service.DeletePatientQuiz(patientHasQuizRemoved)
			if err != nil {
				utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Erro interno! Não foi possível remover o paciente do questionário.", "")
				return
			}
		}
	}
}

func RemoveProximityHasQuizzes(proximitysHasQuizRemoved []models.ProximityHasQuiz, w http.ResponseWriter) {
	for _, proximityHasQuizRemoved := range proximitysHasQuizRemoved {
		alreadyExist, proximityHasQuizRemovedFull := VerifyIfProximityHasQuizExist(proximityHasQuizRemoved)

		if alreadyExist {
			proximityHasQuizRemoved = proximityHasQuizRemovedFull[0]
			err := service.DeleteProximityQuiz(proximityHasQuizRemoved)
			if err != nil {
				utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Erro interno! Não foi possível remover o paciente do questionário.", "")
				return
			}
		}
	}
}

func PutFiliateds(w http.ResponseWriter, r *http.Request) {

	var filiatedsReceived FiliatedsReceived
	var filiatedsRemoved []models.Filiated
	var filiateds []models.Filiated

	err := json.NewDecoder(r.Body).Decode(&filiatedsReceived)
	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusBadRequest, err.Error(), "")
		return
	}
	filiateds = filiatedsReceived.Filiateds
	filiatedsRemoved = filiatedsReceived.FiliatedsRemoved

	patientsHasQuiz, proximitysHasQuiz, _ := ConvertFiliatedsToPatientOrPersonQuizzes(filiateds)

	if len(patientsHasQuiz) != 0 {
		UpdatePatientHasQuizzes(patientsHasQuiz, w)
	}
	if len(proximitysHasQuiz) != 0 {
		UpdateProximityHasQuizzes(proximitysHasQuiz, w)
	}

	patientsHasQuizRemoved, proximitysHasQuizRemoved, _ := ConvertFiliatedsToPatientOrPersonQuizzes(filiatedsRemoved)

	if len(patientsHasQuizRemoved) != 0 {
		RemovePatientHasQuizzes(patientsHasQuizRemoved, w)
	}
	if len(proximitysHasQuizRemoved) != 0 {
		RemoveProximityHasQuizzes(proximitysHasQuizRemoved, w)
	}

	if err != nil {
		utils.ReturnResponseJSON(w, http.StatusInternalServerError, "Erro interno do servidor!", "")
		return
	}

	utils.ReturnResponseJSON(w, http.StatusOK, "Filiados atualizados com sucesso!", "")
	return

}
