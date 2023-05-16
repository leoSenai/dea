package main

import (
	"api/configs"
	"api/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		log.Println("Cannot load configuration from viper")
		panic(err)
	}

	r := chi.NewRouter()

	r.Get("/person/get-by-id/{id}", controller.GetPersonById)
	r.Get("/person/get-all", controller.GetAllPerson)
	r.Post("/person/insert", controller.PostPerson)
	r.Put("/person/update", controller.PutPerson)

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

	err = http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
	if err != nil {
		log.Println("Server not initialized")
		panic(err)
	}
}
