package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/okeyonyia123/gowash/authenticate"
	"github.com/okeyonyia123/gowash/models"
	"github.com/okeyonyia123/gowash/util"
	"github.com/okeyonyia123/gowash/validation"
)

type LoginForm struct {
	PageTitle  string
	FieldNames []string
	Fields     map[string]string
	Errors     map[string]string
}

func DisplayLoginForm(w http.ResponseWriter, r *http.Request, l *LoginForm) {
	fmt.Println("reached display login form")
	PushPage(w, "./views/login.html", nil)

}

func PopulateLoginFormFields(r *http.Request, l *LoginForm) {

	for _, fieldName := range l.FieldNames {
		l.Fields[fieldName] = r.FormValue(fieldName)
	}

}

func ProcessLoginForm(w http.ResponseWriter, r *http.Request, l *LoginForm) {
	var authResult bool
	if "partner" == r.FormValue("type") {
		fmt.Println("a partner is Logging In")
		authResult = authenticate.VerifyPatnerCredentials(r.FormValue("username"), r.FormValue("password"))
		fmt.Println("auth result: ", authResult)

	} else {
		fmt.Printf("a customer is Logging In")
		authResult = authenticate.VerifyCredentials(r.FormValue("username"), r.FormValue("password"))
		fmt.Println("auth result: ", authResult)
	}

	verifyAuth := authResult

	fmt.Println("verifyAuth : ", verifyAuth)

	// Successful login, let's create a cookie for the user and redirect them to the request route
	if verifyAuth == true {

		sessionID := util.GenerateUUID()
		fmt.Println("sessid: ", sessionID)

		//Get connection to the Database
		db, er := validation.GetConnection()
		if er != nil {
			log.Println("Error Connecting to the Database: ", er)
		}

		var u interface{}
		var err error
		if "user" == r.FormValue("type") {
			u, err = db.GetUser(r.FormValue("username"))
			if err != nil {
				log.Print("Encountered error when attempting to fetch user record: ", err)
				http.Redirect(w, r, "/login", 302)
				return
			}

		} else {
			u, err = db.GetPartner(r.FormValue("username"))
			if err != nil {
				log.Print("Encountered error when attempting to fetch user record: ", err)
				http.Redirect(w, r, "/login", 302)
				return
			}

		}

		//Create a Session for that user
		if r.FormValue("type") == "partner" {
			err = authenticate.CreatePartnerSession(u.(*models.Partner), sessionID, w, r)
		} else {
			err = authenticate.CreateUserSession(u.(*models.User), sessionID, w, r)
		}

		if err != nil {
			log.Print("Encountered error when attempting to create user session: ", err)
			http.Redirect(w, r, "/login", 302)
			return

		}

		//Let us decide what page to serve
		var addrs string

		if "user" == r.FormValue("type") {
			uname := l.Fields["username"]
			addrs = "/request/" + uname
		} else {
			uname := l.Fields["username"]
			addrs = "/partner/" + uname
		}

		//RenderTemplate(w, "./views/request.html", l)
		http.Redirect(w, r, addrs, 302)

	} else {

		l.Errors["usernameError"] = "Invalid login."
		DisplayLoginForm(w, r, l)

	}

}

func ValidateLoginForm(w http.ResponseWriter, r *http.Request, l *LoginForm) {

	PopulateLoginFormFields(r, l)
	// Check if username was filled out
	if r.FormValue("username") == "" {
		l.Errors["usernameError"] = "The username field is required."
	}

	// Check if e-mail address was filled out
	if r.FormValue("password") == "" {
		l.Errors["passwordError"] = "The password field is required."
	}

	// Check username syntax
	if validation.CheckUsernameSyntax(r.FormValue("username")) == false {

		usernameErrorMessage := "The username entered has an improper syntax."
		if _, ok := l.Errors["usernameError"]; ok {
			l.Errors["usernameError"] += " " + usernameErrorMessage
		} else {
			l.Errors["usernameError"] = usernameErrorMessage
		}
	}

	if len(l.Errors) > 0 {
		DisplayLoginForm(w, r, l)
	} else {
		fmt.Println("ProcessLoginForm")
		ProcessLoginForm(w, r, l)
	}

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	l := LoginForm{}
	l.FieldNames = []string{"username"}
	l.Fields = make(map[string]string)
	l.Errors = make(map[string]string)
	l.PageTitle = "Log In"

	switch r.Method {

	case "GET":
		DisplayLoginForm(w, r, &l)
	case "POST":
		ValidateLoginForm(w, r, &l)
	default:
		DisplayLoginForm(w, r, &l)
	}
}
