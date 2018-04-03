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
	r.HandleFunc("/signup", handlers.Signup).Methods("GET", "POST")

	//compile mux router
	http.Handle("/", r)

	//Handle favicon request until we have a favicon to work with
	http.Handle("/favicon.ico", http.NotFoundHandler())

	//styling
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./views/gowash/css"))))

	http.ListenAndServe(WEBSERVERPORT, nil)
}
