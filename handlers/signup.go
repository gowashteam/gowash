package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/okeyonyia123/gowash/models"
	"github.com/okeyonyia123/gowash/validation"
)

type Form struct {
	FieldNames  []string
	FieldValues map[string]string
	Errors      map[string]interface{}
}

var validate *validation.Validate

func Signup(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		PushPage(w, "./views/signup.html", Form{})
		return
	}

	//w.Header().Set("content-type", "application/json") //Set content type
	fmt.Println("Signup Handler hit")
	fmt.Println(r.FormValue("email"))

	//Get an intance of a form
	newForm := Form{}
	newForm.FieldNames = []string{"username", "firstname", "lastname", "email", "password", "confirmpassword"} //object literal
	newForm.FieldValues = make(map[string]string)                                                              //initialize a new empty map
	newForm.Errors = make(map[string]interface{})
	//populate FormField with client request so we can work with the form
	PopulateForm(r, &newForm)

	//validate the form inputs
	newForm.Errors = validate.ValidateForm(newForm.FieldValues)

	//check to see if there is any error in the validation process then handle it
	if len(newForm.Errors) > 0 {
		fmt.Println("an error occurered")
		log.Println(newForm.Errors)
		//fmt.Println(newForm.Errors)
		//w.WriteHeader(http.StatusNotAcceptable)
		//json.NewEncoder(w).Encode(newForm)
		//PushPage(w, "./views/login.html", newForm)
		DisplaySignUpForm(w, r, &newForm)
		return
	}

	if r.FormValue("type") == "partner" {
		//pass the form to a function that builds a Partner Model
		partner := models.NewPartner(newForm.FieldValues)
		//Pass the Model to the function that talks to the database
		postedPartner, err := validate.QuerryDB(partner, "postPartner")

		//If everything goes well, return the model back to the user
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		fmt.Println(postedPartner)
		PushPage(w, "./views/signupconfirmation.html", newForm)

	} else {
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
		PushPage(w, "./views/signupconfirmation.html", newForm)
	}

	//fmt.Fprint(w, postedUser) //actual response written to client
	//json.NewEncoder(w).Encode(postedUser)
	w.WriteHeader(http.StatusOK) //Set the header

}

func DisplaySignUpForm(w http.ResponseWriter, r *http.Request, s *Form) {
	//tmpl.ExecuteTemplate(w, "signupform.html", s)
	PushPage(w, "./views/signup.html", s)

}

func PopulateForm(r *http.Request, f *Form) {
	//populate the formFields map
	for _, field := range f.FieldNames {
		f.FieldValues[field] = r.FormValue(field)
	}
}
