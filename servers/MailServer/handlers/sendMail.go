package handlers

import (
	"net/http"
)

func SendMail(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mail Server Up and running at port 8082"))
}
