package main

import (
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

var templates *template.Template

func RouteHello(w http.ResponseWriter, r *http.Request) {
	itallo := User{"Itallo", "male", 20}
	
	templates.ExecuteTemplate(w, "hello.html", itallo)
}

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", RouteHello)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
