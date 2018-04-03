package handlers

import (
	"net/http"
)

func ContactUsHandler(w http.ResponseWriter, r *http.Request) {
	PushPage(w, "./views/contact-us.html", nil)
}
