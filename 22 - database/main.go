package main

import (
	"LearnDB/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/users", routes.CreateUser).Methods("POST")
	router.HandleFunc("/users", routes.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", routes.UpdateUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":5000", router))
}
