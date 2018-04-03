package handlers

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	PushPage(w, "./views/index.html", nil)
}
