package main

import (
	"net/http"

	"github.com/michalslomczynski/shortest-ways/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/routes", handler.HandleRequest).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8086", nil)
}
