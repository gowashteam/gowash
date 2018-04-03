package handlers

import (
	"net/http"
)

func ServicesUsHandler(w http.ResponseWriter, r *http.Request) {
	PushPage(w, "./views/services.html", nil)
}
