package main

import (
	"log"
	"net/http"

	"github.com/michalslomczynski/shortest-ways/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/routes", handler.HandleRequest).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8086", nil))
}
