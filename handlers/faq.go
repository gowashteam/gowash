package handlers

import (
	"net/http"
)

func FaqHandler(w http.ResponseWriter, r *http.Request) {
	PushPage(w, "./views/faq.html", nil)
}
