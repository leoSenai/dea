package main

import (
	"api/configs"
	"api/controller"
	"api/tests"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {

	err := configs.Load()
	if err != nil {
		log.Println("Cannot load configuration from viper")
		panic(err)
	} else {
		conf := configs.GetDB()
		fmt.Println("[INFO] Database host is on " + conf.Host)
		fmt.Println("[INFO] Listening...")
	}

	tests.Run()

	r := chi.NewRouter()

	// Configurar o middleware de CORS para permitir acesso de qualquer origem
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Permitir acesso de qualquer origem (isso é apenas para desenvolvimento, em produção, especifique as origens permitidas)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	r.Use(corsHandler.Handler) // Use o middleware de CORS em todas as rotas

	r.Post("/auth/login", controller.PostLogin)

	r.Get("/person/get-by-id/{id}", controller.GetPersonById)
	r.Get("/person/get-all", controller.GetAllPerson)
	r.Post("/person/insert", controller.PostPerson)
	r.Put("/person/update", controller.PutPerson)
	r.Get("/person/get-by-doc/{docNumber}", controller.GetPersonByDocNumber)

	r.Get("/user/get-by-id/{id}", controller.GetUserById)
	r.Get("/user/get-all", controller.GetAllUser)
	r.Post("/user/insert", controller.PostUser)
	r.Put("/user/update", controller.PutUser)

	r.Get("/cbo/get-by-id/{id}", controller.GetCboById)
	r.Get("/cbo/get-all", controller.GetAllCbo)
	r.Post("/cbo/insert", controller.PostCbo)
	r.Put("/cbo/update", controller.PutCbo)

	r.Get("/patient/get-by-id/{id}", controller.GetPatientById)
	r.Get("/patient/get-all", controller.GetAllPatient)
	r.Post("/patient/insert", controller.PostPatient)
	r.Put("/patient/update", controller.PutPatient)

	r.Get("/doctor/get-by-id/{id}", controller.GetDoctorById)
	r.Get("/doctor/get-all", controller.GetAllDoctor)
	r.Post("/doctor/insert", controller.PostDoctor)
	r.Put("/doctor/update", controller.PutDoctor)

	r.Get("/services/get-by-id/{id}", controller.GetServicesById)
	r.Get("/services/get-all", controller.GetAllServices)
	r.Post("/services/insert", controller.PostServices)
	r.Put("/services/update", controller.PutServices)

	r.Get("/quiz/get-by-id/{id}", controller.GetQuizById)
	r.Get("/quiz/get-all", controller.GetAllQuiz)
	r.Post("/quiz/insert", controller.PostQuiz)
	r.Put("/quiz/update", controller.PutQuiz)

	r.Get("/question/get-by-id/{id}", controller.GetQuestionById)
	r.Get("/question/get-all", controller.GetAllQuestion)
	r.Post("/question/insert", controller.PostQuestion)
	r.Put("/question/update", controller.PutQuestion)
	r.Put("/question/update-bulk", controller.PutQuestionsBulk)
	r.Put("/question/get-by-quiz/{idQuiz}", controller.GetQuestionsByQuiz)
	r.Delete("/questions/{id}", controller.DeleteQuestionById)

	r.Get("/anamnese/get-by-id/{id}", controller.GetAnamneseById)
	r.Get("/anamnese/get-all", controller.GetAllAnamnese)
	r.Post("/anamnese/insert", controller.PostAnamnese)
	r.Put("/anamnese/update", controller.PutAnamnese)

	r.Get("/asking/get-by-id/{id}", controller.GetAskingById)
	r.Get("/asking/get-all", controller.GetAllAsking)
	r.Post("/asking/insert", controller.PostAsking)
	r.Put("/asking/update", controller.PutAsking)

	r.Post("/proximity/insert", controller.PostProximity)
	r.Get("/proximity/get-by-id-person/{id}", controller.GetProximityAllByIdPerson)
	r.Get("/proximity/get-by-id-patient/{id}", controller.GetProximityAllByIdPatient)

	r.Get("/anamnesehasasking/get-by-anamnese-id/{id}", controller.GetAnamneseHasAskingByAnamneseId)
	r.Get("/anamnesehasasking/get-by-asking-id/{id}", controller.GetAnamneseHasAskingByAskingId)
	r.Get("/anamnesehasasking/get-all", controller.GetAllAnamneseHasAsking)
	r.Post("/anamnesehasasking/insert", controller.PostAnamneseHasAsking)

	r.Get("/proximityhasquiz/get", controller.GetProximityQuizByQuizPatientPersonIDs)
	r.Get("/proximityhasquiz/get-by-id-quiz/{id}", controller.GetProximityQuizByQuizID)
	r.Get("/proximityhasquiz/get-by-id-patient/{id}", controller.GetProximityQuizByPatientID)
	r.Get("/proximityhasquiz/get-by-id-person/{id}", controller.GetProximityQuizByPersonID)
	r.Post("/proximityhasquiz/insert", controller.PostProximityQuiz)
	r.Post("/proximityhasquiz/update", controller.PutProximityQuiz)

	r.Get("/patienthasdoctor/get-by-patient-id/{id}", controller.GetPatientHasDoctorByPatientId)
	r.Get("/patienthasdoctor/get-by-doctor-id/{id}", controller.GetPatientHasDoctorByDoctorId)
	r.Get("/patienthasdoctor/get-all", controller.GetAllPatientHasDoctor)
	r.Post("/patienthasdoctor/insert", controller.PostPatientHasDoctor)

	err = http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
	if err != nil {
		log.Println("Server not initialized")
		panic(err)
	}
}
