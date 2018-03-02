package handlers

import (
	"fmt"
	"net/http"

	"github.com/okeyonyia123/gowash/servers/Signup/models"
	"github.com/okeyonyia123/gowash/servers/Signup/validation"
)

type Form struct {
	FieldNames  []string
	FieldValues map[string]string
	Errors      map[string]string
}

func Signup(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) //Set the header
	//w.Write([]byte("Signup handler is Up"))
	fmt.Println("Signup Handler hit")

	//Get an intance of a form
	newForm := form{}
	newForm.FieldNames = []string{"username","firstname", "lastname" "email", "password"} //object literal
	newForm.FieldValues = make(map[string]string)      //initialize a new empty map
	newForm.Errors = make(map[string]string)           //initialize a new map for error map

	//send the request body to a validation function
	PopulateForm(r, &newForm) //the form is populated and ready to be returned to the client

	//validate the form inputs
	//validate Email
	validation.validateForm(&newForm)

	//check to see if form.errors isempty
	if len(newForm.Errors) > 0 {
		fmt.Fprint(w, newForm)
		return
	}

	//pass the form to a function that builds a user Model
	user := models.NewUser(newForm)

	//Pass the Model to the function that talks to the database
	postedUser, err := validation.querryDB(user, postUser)

	//If everything goes well, return the model back to the user
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, postedUser)

}

func PopulateForm(r *http.Request, f *Form) {
	//populate the formFields map
	for _, field := range f.FieldNames {
		f.FieldValues[field] = r.FormValue(field)
	}
}
