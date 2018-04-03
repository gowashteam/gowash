package handlers

import (
	"net/http"
)

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	PushPage(w, "./views/forgot-password.html", nil)
}
