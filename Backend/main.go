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

	r.Get("/person/get-by-id", controller.GetPersonById)

	r.Get("/user/get-by-id", controller.GetUserById)
	r.Get("/user/get-all", controller.GetAllUsers)
	r.Post("/user/insert", controller.PostUser)
	r.Put("/user/update", controller.UpdateUser)

	r.Get("/cbo/get-by-id", controller.GetCboById)
	r.Get("/cbo/get-all", controller.GetAllCbos)
	r.Post("/cbo/insert", controller.PostCbo)
	r.Put("/cbo/update", controller.UpdateCbo)

	err = http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
	if err != nil {
		log.Println("Server not initialized")
		panic(err)
	}
}
