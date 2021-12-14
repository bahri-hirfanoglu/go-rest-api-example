package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouter() {

	r := mux.NewRouter()

	r.HandleFunc("/players", GetPlayers).Methods("GET")
	r.HandleFunc("/player/{id}", GetPlayer).Methods("GET")
	r.HandleFunc("/player", AddPlayer).Methods("POST")
	r.HandleFunc("/player/{id}", DeletePlayer).Methods("DELETE")
	r.HandleFunc("/player/{id}", UpdatePlayer).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8088", r))
}
