package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/okeyonyia123/gowash/servers/Login/handlers"
)

func main() {
	fmt.Println("Login server is up and running")

	//get a new intance of mux
	r := mux.NewRouter()
	r.HandleFunc("/login", handlers.Login).Methods("POST")

}
