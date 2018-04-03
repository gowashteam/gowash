package main

import (
	"github.com/gorilla/mux"
	"github.com/okeyonyia123/gowash/servers/MailServer/handlers"
)

func main() {

	//get a routter
	r := mux.NewRouter()
	r.HandleFunc("/mail", handlers.SendMail).Methods("POST", "GET")
}
