package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func PushPage(w http.ResponseWriter, filePath string, data interface{}) {
	// Get a template
	t, err := template.ParseFiles(filePath)
	if err != nil {
		fmt.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Execute Templae
	t.Execute(w, data)
}
