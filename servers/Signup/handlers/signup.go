package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/okeyonyia123/gowash/servers/Signup/models"
	"github.com/okeyonyia123/gowash/servers/Signup/validation"
)

type Form struct {
	FieldNames  []string
	FieldValues map[string]string
	Errors      map[string]error
}

var validate *validation.Validate

func Signup(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json") //Set content type
	fmt.Println("Signup Handler hit")
	fmt.Println(r.FormValue("email"))

	//Get an intance of a form
	newForm := Form{}
	newForm.FieldNames = []string{"username", "firstname", "lastname", "email", "password", "confirmpassword"} //object literal
	newForm.FieldValues = make(map[string]string)                                                              //initialize a new empty map
	newForm.Errors = make(map[string]error)
	//populate FormField with client request so we can work with the form
	PopulateForm(r, &newForm)

	//validate the form inputs
	newForm.Errors = validate.ValidateForm(newForm.FieldValues)

	//check to see if there is any error in the validation process then handle it
	if len(newForm.Errors) > 0 {
		fmt.Println("an error occurered")
		fmt.Println(newForm.Errors)
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(newForm)
		return
	}

	//pass the form to a function that builds a user Model
	user := models.NewUser(newForm.FieldValues)

	//Pass the Model to the function that talks to the database
	postedUser, err := validate.QuerryDB(user, "postUser")

	//If everything goes well, return the model back to the user
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Println(postedUser)

	//fmt.Fprint(w, postedUser) //actual response written to client
	json.NewEncoder(w).Encode(postedUser)
	w.WriteHeader(http.StatusOK) //Set the header

}

func PopulateForm(r *http.Request, f *Form) {
	//populate the formFields map
	for _, field := range f.FieldNames {
		f.FieldValues[field] = r.FormValue(field)
	}
}
