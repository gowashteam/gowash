package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/okeyonyia123/gowash/servers/Signup/handlers"
)

const (
	WEBSERVERPORT = ":8084"
)

func main() {
	//Establish Connection with database

	r := mux.NewRouter()
	//Homepage router
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	//signup handler
	r.HandleFunc("/signup", handlers.Signup).Methods("GET")

	//login
	r.HandleFunc("/login", handlers.Login).Methods("GET")

	//compile mux router
	http.Handle("/", r)

	http.ListenAndServe(WEBSERVERPORT, nil)
}
