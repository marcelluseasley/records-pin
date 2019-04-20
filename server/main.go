package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	_ "github.com/mattn/go-sqlite3"
)

// Record represents call data that will be stored in the database
type Record struct {
	PhoneNumber string  `json:"phone_number"`
	Carrier     string  `json:"carrier"`
	Score       float64 `json:"score"`
}

func main() {
	createdatabase()

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/v1/record", func(r chi.Router) {
		r.Post("/create", createRecordHandler)
		r.Get("/read/{phoneNumber}", readRecordHandler)
		r.Put("/update/{phoneNumber}", updateRecordHandler)
		r.Delete("/delete", deleteRecordHandler) //TODO: add authentication
	})

	router.Route("/v1/records", func(r chi.Router) {
		r.Post("/listall", listAllHandler)
		r.Delete("/deleteall", deleteAllHandler) //TODO: add authentication

	})

	port := "8080"
	log.Printf("Starting server on port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	
}
