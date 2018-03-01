package handlers

import (
	"fmt"
	"net/http"
)

type form struct {
	fieldNames  []string
	fieldValues map[string]string
	erros       map[string]string
}

func Signup(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) //Set the header
	//w.Write([]byte("Signup handler is Up"))
	fmt.Println("Signup Handler hit")
	//initialize the form fields
	newForm := form{}
	newForm.fieldNames = []string{"username", "email", "password", "confirmpassword"} //object literal
	newForm.fieldValues = make(map[string]string)                                     //initialize a new empty map
	newForm.erros = make(map[string]string)                                           //initialize a new map for error map
	//send the request body to a validation function
	//Get an intance of a form

	//validate the form inputs

	//pass the form to a function that builds a user Model

	//Pass the Model to the function that talks to the database

	//If everything goes well, return the model back to the user

}
