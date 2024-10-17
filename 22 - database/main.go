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

	log.Fatal(http.ListenAndServe(":5000", router))
}
