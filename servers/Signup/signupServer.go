package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/okeyonyia123/gowash/servers/Signup/handlers"
	"github.com/okeyonyia123/gowash/servers/signup/datastore"
)

const (
	WEBSERVERPORT = ":8084"
)

func main() {
	//Establish Connection with database
	db, err := datastore.NewDatastore(datastore.MONGODB, "159.65.188.249:27017")

	if err != nil {
		log.Print(err)
	} else {
		fmt.Println(db) //log the connection to indicate successful connection
	}

	defer db.Close() //close connection

	r := mux.NewRouter()
	//Homepage router
	r.HandleFunc("/", handlers.HomeHandler)

	//compile mux router
	http.Handle("/", r)

	http.ListenAndServe(WEBSERVERPORT, nil)
}
