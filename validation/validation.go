package validation

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/okeyonyia123/gowash/datastore"
	"github.com/okeyonyia123/gowash/models"
)

type Validate struct {
	Errors     map[string]interface{}
	FormFields map[string]string
}

const UsernameRegex string = `^@?(\w){1,15}$`
const EmailRegex = `(?i)^[_a-z0-9-]+(\.[_a-z0-9-]+)*@[a-z0-9-]+(\.[a-z0-9-]+)*(\.[a-z]{2,3})+$`

func CheckUsernameSyntax(username string) bool {

	r, err := regexp.Compile(UsernameRegex)
	if err != nil {
		log.Fatal(err)
	}

	return r.MatchString(username)
}

func CheckEmailSyntax(email string) bool {

	r, err := regexp.Compile(EmailRegex)
	if err != nil {
		log.Fatal(err)
	}

	return r.MatchString(email)
}

func GetConnection() (datastore.Datastore, error) {
	return datastore.NewDatastore(datastore.MONGODB, "159.65.188.249:27017")
}

func (v *Validate) QuerryDB(querry interface{}, querryType string) (interface{}, error) {
	db, err := GetConnection()

	var person interface{}

	if err != nil {
		fmt.Println("Error Here")
		log.Print(err)

	} else {
		fmt.Printf("Established Connection to Database on : %v", db) //log the connection to indicate successful connection
		fmt.Println()
	}

	//Switch querryType to talk to the database depending on what exactly we want
	switch querryType {
	case "getUser":
		aUser, err := db.GetUser(querry.(string)) //type assertion to convert the interface to actual type
		if err != nil {
			return nil, err
		}
		person = aUser

	case "postUser":
		aUser, err := db.PostUser(querry.(*models.User)) //type assertion to convert the interface to actual type
		if err != nil {
			return nil, err
		}
		person = aUser

	case "postPartner":
		aPartner, err := db.PostPartner(querry.(*models.Partner)) //type assertion to convert the interface to actual type
		if err != nil {
			return nil, err
		}
		person = aPartner

	}

	defer db.Close() //close connection

	return person, nil

}

func (v *Validate) ValidateEmail(email string) (bool, error) {
	//check if valid email
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	matched := re.MatchString(email)
	if !matched {
		return false, errors.New("Please enter a valid Email")
	}

	//check if valid email already exists on database
	user, _ := v.QuerryDB(email, "getUSer")
	if user != nil {
		//email already exixts in database
		return false, errors.New("Email already exixts")
	}
	fmt.Println(user)
	return true, nil
}

func (v *Validate) ValidateUsername(username string, FieldValues map[string]string) (bool, error) {
	if username == "" {
		return false, errors.New("username field cannot be empty")
	}

	//check to see if username is taken
	user, _ := v.QuerryDB(username, "getUser")
	if user != nil {
		return false, errors.New("username is taken")
	}

	return true, nil

}

func (v *Validate) ValidateFirstName(firstname string) (bool, error) {

	if firstname == "" {
		return false, errors.New("firstname field cannot be empty")
	}

	return true, nil
}

func (v *Validate) ValidateLastName(lastname string) (bool, error) {

	if lastname == "" {
		return false, errors.New("lastname field cannot be empty")
	}

	return true, nil
}

func (v *Validate) ValidatePassword(password string) (bool, error) {

	if password == "" {
		return false, errors.New("password field cannot be empty")
	}

	return true, nil
}

func (v *Validate) ValidateConfirmPassword(FieldValues map[string]string) (bool, error) {

	if len(FieldValues["confirmpassword"]) == 0 {
		return false, errors.New("confirmpassword field cannot be empty")
	}

	if FieldValues["confirmpassword"] != FieldValues["password"] {
		return false, errors.New("confirmpassword field must match password field")
	}

	return true, nil
}

func (validator *Validate) ValidateForm(FieldValues map[string]string) map[string]interface{} {
	//initialize error map to store new errors
	v := Validate{}
	v.Errors = make(map[string]interface{})

	//range over the form and validate the various fields
	for field, value := range FieldValues {
		switch field {

		case "email":
			if isValid, err := v.ValidateEmail(value); !isValid {
				v.Errors[field] = err
			}
			break

		case "username":
			if isValid, err := v.ValidateUsername(value, FieldValues); !isValid {
				v.Errors[field] = err
			}
			break

		case "firstname":
			if isValid, err := v.ValidateFirstName(value); !isValid {
				v.Errors[field] = err
			}
			break

		case "lastname":
			if isValid, err := v.ValidateLastName(value); !isValid {
				v.Errors[field] = err
			}
			break

		case "password":
			if isValid, err := v.ValidatePassword(value); !isValid {
				v.Errors[field] = err
			}
			break

		case "confirmpassword":
			if isValid, err := v.ValidateConfirmPassword(FieldValues); !isValid {
				v.Errors[field] = err
			}
			break

		default:
			v.Errors["default"] = errors.New("cannot validate form fields provided. Please refer to the API documentation for possible form fields")
			return v.Errors
		}
	}

	return v.Errors
}
