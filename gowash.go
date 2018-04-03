package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/okeyonyia123/gowash/handlers"
)

const (
	WEBSERVERPORT = ":8084"
)

func main() {
	//Establish Connection with database

	r := mux.NewRouter()

	//Handle favicon request we have a favicon to work with
	http.Handle("/favicon.ico", http.NotFoundHandler())

	//Homepage router
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	//signup handler
	r.HandleFunc("/signup", handlers.Signup).Methods("GET", "POST")
	//Login Handler
	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")

	//Welcome Page handler
	r.HandleFunc("/partner/{username}", handlers.PartnerHandler).Methods("GET", "POST")

	//compile mux router
	http.Handle("/", r)
	//static routes
	r.HandleFunc("/about-us", handlers.AboutUsHandler).Methods("GET", "POST")
	r.HandleFunc("/services", handlers.ServicesUsHandler).Methods("GET", "POST")
	r.HandleFunc("/contact-us", handlers.ContactUsHandler).Methods("GET", "POST")
	r.HandleFunc("/faq", handlers.FaqHandler).Methods("GET", "POST")
	r.HandleFunc("/forgot-password", handlers.ForgotPasswordHandler).Methods("GET", "POST")

	//styling
	r.PathPrefix("/views/").Handler(http.StripPrefix("/views/", http.FileServer(http.Dir("./views"))))
	//r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./views/gowash/css"))))

	http.ListenAndServe(WEBSERVERPORT, nil)
}
