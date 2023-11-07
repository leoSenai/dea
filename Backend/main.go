package main

import (
	"api/configs"
	"api/controller"
	"api/tests"
	"api/utils"
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

	r.Get("/person/get-by-id/{id}", utils.VerifyToken(controller.GetPersonById))
	r.Get("/person/get-all", utils.VerifyToken(controller.GetAllPerson))
	r.Post("/person/insert", utils.VerifyToken(controller.PostPerson))
	r.Put("/person/update", utils.VerifyToken(controller.PutPerson))
	r.Get("/person/get-by-doc/{docNumber}", utils.VerifyToken(controller.GetPersonByDocNumber))
	r.Put("/person/reset-password", utils.VerifyToken(controller.ResetPassword))

	r.Get("/user/get-by-id/{id}", utils.VerifyToken(controller.GetUserById))
	r.Get("/user/get-all", utils.VerifyToken(controller.GetAllUser))
	r.Post("/user/insert", utils.VerifyToken(controller.PostUser))
	r.Put("/user/update", utils.VerifyToken(controller.PutUser))
	r.Put("/user/reset-password", utils.VerifyToken(controller.ResetPasswordUser))

	r.Get("/cbo/get-by-id/{id}", utils.VerifyToken(controller.GetCboById))
	r.Get("/cbo/get-all", utils.VerifyToken(controller.GetAllCbo))
	r.Post("/cbo/insert", utils.VerifyToken(controller.PostCbo))
	r.Put("/cbo/update", utils.VerifyToken(controller.PutCbo))

	r.Get("/patient/get-by-id/{id}", utils.VerifyToken(controller.GetPatientById))
	r.Get("/patient/get-all", utils.VerifyToken(controller.GetAllPatient))
	r.Post("/patient/insert", utils.VerifyToken(controller.PostPatient))
	r.Put("/patient/update", utils.VerifyToken(controller.PutPatient))
	r.Put("/patient/reset-password", utils.VerifyToken(controller.ResetPasswordPatient))

	r.Get("/doctor/get-by-id/{id}", utils.VerifyToken(controller.GetDoctorById))
	r.Get("/doctor/get-all", utils.VerifyToken(controller.GetAllDoctor))
	r.Post("/doctor/insert", utils.VerifyToken(controller.PostDoctor))
	r.Put("/doctor/update", utils.VerifyToken(controller.PutDoctor))

	r.Get("/services/get-by-id/{id}", utils.VerifyToken(controller.GetServicesById))
	r.Get("/services/get-all", utils.VerifyToken(controller.GetAllServices))
	r.Post("/services/insert", utils.VerifyToken(controller.PostServices))
	r.Put("/services/update", utils.VerifyToken(controller.PutServices))

	r.Get("/quiz/get-by-id/{id}", utils.VerifyToken(controller.GetQuizById))
	r.Get("/quiz/get-all", utils.VerifyToken(controller.GetAllQuiz))
	r.Post("/quiz/insert", utils.VerifyToken(controller.PostQuiz))
	r.Put("/quiz/update", utils.VerifyToken(controller.PutQuiz))

	r.Get("/question/get-by-id/{id}", utils.VerifyToken(controller.GetQuestionById))
	r.Get("/question/get-all", utils.VerifyToken(controller.GetAllQuestion))
	r.Post("/question/insert", utils.VerifyToken(controller.PostQuestion))
	r.Put("/question/update", utils.VerifyToken(controller.PutQuestion))
	r.Put("/question/update-bulk", utils.VerifyToken(controller.PutQuestionsBulk))
	r.Put("/question/get-by-quiz/{idQuiz}", utils.VerifyToken(controller.GetQuestionsByQuiz))
	r.Delete("/question/delete/{id}", utils.VerifyToken(controller.DeleteQuestionById))

	r.Get("/anamnese/get-by-id/{id}", utils.VerifyToken(controller.GetAnamneseById))
	r.Get("/anamnese/get-by-id-user-patient/{idUser}/{idPatient}", utils.VerifyToken(controller.GetAnamneseByIdUserPatient))
	r.Get("/anamnese/get-all", utils.VerifyToken(controller.GetAllAnamnese))
	r.Post("/anamnese/insert", utils.VerifyToken(controller.PostAnamnese))
	r.Put("/anamnese/update", utils.VerifyToken(controller.PutAnamnese))

	r.Get("/asking/get-by-id/{id}", utils.VerifyToken(controller.GetAskingById))
	r.Get("/asking/get-all", utils.VerifyToken(controller.GetAllAsking))
	r.Post("/asking/insert", utils.VerifyToken(controller.PostAsking))
	r.Put("/asking/update", utils.VerifyToken(controller.PutAsking))

	r.Post("/proximity/insert", utils.VerifyToken(controller.PostProximity))
	r.Get("/proximity/get-by-id-person/{id}", utils.VerifyToken(controller.GetProximityAllByIdPerson))
	r.Get("/proximity/get-by-id-patient/{id}", utils.VerifyToken(controller.GetProximityAllByIdPatient))
	r.Get("/proximity/get-persons-by-id-patient/{id}", utils.VerifyToken(controller.GetPersonNoPasswordProximityAllByIdPatient))

	r.Get("/anamnesehasasking/get-by-anamnese-id/{id}", utils.VerifyToken(controller.GetAnamneseHasAskingByAnamneseId))
	r.Get("/anamnesehasasking/get-by-asking-id/{id}", utils.VerifyToken(controller.GetAnamneseHasAskingByAskingId))
	r.Get("/anamnesehasasking/get-all", utils.VerifyToken(controller.GetAllAnamneseHasAsking))
	r.Post("/anamnesehasasking/insert", utils.VerifyToken(controller.PostAnamneseHasAsking))

	r.Get("/proximityhasquiz/get-all", utils.VerifyToken(controller.GetProximityQuizByQuizPatientPersonIDs))
	r.Get("/proximityhasquiz/get-by-id-quiz/{id}", utils.VerifyToken(controller.GetProximityQuizByQuizID))
	r.Get("/proximityhasquiz/get-by-id-patient/{id}", utils.VerifyToken(controller.GetProximityQuizByPatientID))
	r.Get("/proximityhasquiz/get-by-id-person/{id}", utils.VerifyToken(controller.GetProximityQuizByPersonID))
	r.Get("/proximityhasquiz/get-by-id-quiz-person/{idquiz}/{idperson}", utils.VerifyToken(controller.GetProximityQuizByQuizPersonID))
	r.Post("/proximityhasquiz/insert", utils.VerifyToken(controller.PostProximityQuiz))
	r.Put("/proximityhasquiz/update", utils.VerifyToken(controller.PutProximityQuiz))

	r.Get("/patienthasdoctor/get-by-patient-id/{id}", utils.VerifyToken(controller.GetPatientHasDoctorByPatientId))
	r.Get("/patienthasdoctor/get-by-doctor-id/{id}", utils.VerifyToken(controller.GetPatientHasDoctorByDoctorId))
	r.Get("/patienthasdoctor/get-all", utils.VerifyToken(controller.GetAllPatientHasDoctor))
	r.Post("/patienthasdoctor/insert", utils.VerifyToken(controller.PostPatientHasDoctor))

	r.Get("/patienthasquiz/get-by-id-quiz/{id}", utils.VerifyToken(controller.GetPatientQuizByQuizID))
	r.Get("/patienthasquiz/get-by-id-quiz-patient/{idquiz}/{idpatient}", utils.VerifyToken(controller.GetPatientQuizByQuizPatientID))
	r.Get("/patienthasquiz/get-by-id-patient/{id}", utils.VerifyToken(controller.GetPatientQuizByPatientID))
	r.Post("/patienthasquiz/insert", utils.VerifyToken(controller.PostPatientQuiz))
	r.Put("/patienthasquiz/update", utils.VerifyToken(controller.PutPatientQuiz))

	r.Put("/filiateds/update", utils.VerifyToken(controller.PutFiliateds))

	err = http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
	if err != nil {
		log.Println("Server not initialized")
		panic(err)
	}
}
