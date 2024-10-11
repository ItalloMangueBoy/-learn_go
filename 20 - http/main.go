package main

import (
	"log"
	"net/http"
)

func RouteHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/hello", RouteHello	)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
