package handlers

import (
	"net/http"
)

func AboutUsHandler(w http.ResponseWriter, r *http.Request) {
	PushPage(w, "./views/about-us.html", nil)
}
