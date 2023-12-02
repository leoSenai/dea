package repository

import (
	"api/db"
	"api/models"
	"log"
	"strconv"
	"strings"
)

func GetAnamneseById(id int64) (anamnese models.Anamnese, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.First(&anamnese, id)
	log.Printf("row: %v", row)

	return
}

func GetAnamneseByIdUserPatient(idUser int64, idPatient int64) (anamnese models.Anamnese, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.Where("usuario_idusuario = ? AND paciente_idpaciente = ?", idUser, idPatient).Find(&anamnese)
	log.Printf("row: %v", row)

	return
}

func GetAllAnamnese() (anamneses []models.Anamnese, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	rows := conn.Find(&anamneses)
	log.Printf("rows: %v", rows)

	return
}

func PostAnamnese(anamnesePost models.Anamnese) (anamneseBack models.Anamnese, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.Create(&anamnesePost)
	log.Printf("row: %v", row)
	conn.First(&anamneseBack, anamnesePost.IdAnamnese)

	return
}

func PutAnamnese(anamnesePut models.Anamnese) (anamneseBack models.Anamnese, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	anamneseBack = anamnesePut

	if anamnesePut.IdAnamnese != 0 {
		row := conn.Table("anamnese").Where("idanamnese = ?", anamnesePut.IdAnamnese).Updates(&anamneseBack)
		conn.Table("anamnese").Select("indicativo").Where("idanamnese = ?", anamnesePut.IdAnamnese).Updates(&anamneseBack)
		log.Printf("row: %v", row)
		conn.First(&anamneseBack, anamnesePut.IdAnamnese)
	}

	return
}

func GetSumResults(idpatient int) (sumresults []models.SumResult, err error) {
	//sumresults - agrupado de somatorios de cada questionario de pessoa prox e de paciente

	patientHasQuizzes, _ := GetPatientQuizByPatientID(int64(idpatient))
	proximityHasQuizzes, _ := GetProximityQuizByPatientID(int64(idpatient))

	if len(patientHasQuizzes) != 0 {
		for _, patientHasQuiz := range patientHasQuizzes {
			if patientHasQuiz.Finished == 1 {

				patient, _ := GetPatientById(int64(idpatient))
				quiz, _ := GetQuizById(patientHasQuiz.IdQuiz)

				sumAnswersStr := strings.Split(patientHasQuiz.Answers, ";")
				var answersInt []int
				var sumAnswers int

				for _, numStr := range sumAnswersStr {
					num, _ := strconv.Atoi(numStr)
					answersInt = append(answersInt, num)
				}

				for _, num := range answersInt {
					sumAnswers += num
				}

				sumresults = append(sumresults, models.SumResult{
					ProximityName: patient.Name,
					ProximityDesc: patient.Name,
					QuizDesc:      quiz.Name,
					Sum:           sumAnswers,
				})
			}
		}
	}
	if len(proximityHasQuizzes) != 0 {
		for _, proximityHasQuiz := range proximityHasQuizzes {
			if proximityHasQuiz.Finished == 1 {
				person, _ := GetPersonById(int64(idpatient))
				quiz, _ := GetQuizById(proximityHasQuiz.IdQuiz)

				proximitys, _ := GetProximityAllByIdPatient(int64(idpatient))

				var proximityDesc = ""

				for _, prox := range proximitys {
					if prox.IdPerson == person.IdPerson {
						proximityDesc = prox.Desc
					}
				}

				sumAnswersStr := strings.Split(proximityHasQuiz.Answers, ";")
				var answersInt []int
				var sumAnswers int

				for _, numStr := range sumAnswersStr {
					num, _ := strconv.Atoi(numStr)
					answersInt = append(answersInt, num)
				}

				for _, num := range answersInt {
					sumAnswers += num
				}

				sumresults = append(sumresults, models.SumResult{
					ProximityName: person.Name,
					ProximityDesc: proximityDesc,
					QuizDesc:      quiz.Name,
					Sum:           sumAnswers,
				})
			}
		}
	}
	return
}
