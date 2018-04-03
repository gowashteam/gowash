package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var m = make(map[string]string)

func PartnerHandler(w http.ResponseWriter, r *http.Request) {
	m["PageTitle"] = "Partner Profile"
	vars := mux.Vars(r)
	username := vars["username"]
	fmt.Println(username)
	//w.Write([]byte(username))
	m["username"] = username
	fmt.Println(m)
	PushPage(w, "./views/profile.html", m)

}
