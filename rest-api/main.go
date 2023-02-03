package main

import (
	"log"
	"net/http"
	"rest-api/handlers"
	"rest-api/models"

	"github.com/gorilla/mux"
)

func main() {

	DB := models.Connect()
	h := handlers.New(DB)
	r := mux.NewRouter()

	r.HandleFunc("/api/products", h.GetAllBarang).Methods(http.MethodGet)
	r.HandleFunc("/api/products/{id}", h.GetBarang).Methods(http.MethodGet)
	r.HandleFunc("/api/products", h.AddBarang).Methods(http.MethodPost)
	r.HandleFunc("/api/products/{id}", h.UpdateBarang).Methods(http.MethodPut)
	r.HandleFunc("/api/products/{id}", h.DeleteBarang).Methods(http.MethodDelete)

	log.Println("API RUN 8080")
	http.ListenAndServe(":8080", r)

}
