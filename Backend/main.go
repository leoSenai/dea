package main

import (
	"api/configs"
	"api/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Get("/get-by-id", handlers.Get)
	r.Get("/", handlers.GetTeste)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
