package handlers

import (
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) //Set the header
	w.Write([]byte("Login Handler is Up"))
}
