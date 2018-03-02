package validation

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/okeyonyia123/gowash/servers/Signup/datastore"
	"github.com/okeyonyia123/gowash/servers/Signup/handlers"
)

func QuerryDB(interface querry, string querryType) (string, error) {
	db, err := datastore.NewDatastore(datastore.MONGODB, "159.65.188.249:27017")

	if err != nil {
		log.Print(err)
	} else {
		fmt.Println("Established Connection to Database on: %v", db) //log the connection to indicate successful connection
	}

  //Switch querryType to talk to the database depending on what exactly we want
  switch querryType {
  case getUser:
    user, err := db.GetUser(querry)
    if err != nil {
      return nil, err
    }
    return user nil

  case postUser:
    err := db.PostUser(querry)
    if err != nil {
      return nil, err
    }
    return nil,nil

  }

	defer db.Close() //close connection

}

func ValidateEmail(email string) (bool, error) {
	//check if valid email
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	matched := re.MatchString(email)
	if !matched {
		return false, errors.New("Please enter a valid Email")
	}

	//check if valid email already exists on database
	user, _ := QuerryDB(email)

	if user != nil {
		//email already exixts in database
		return false, errors.new("Email already exixts")
	}
	return true, nil
}

func ValidateUsername(username string) (bool, error) {
	if username == "" {
		return false, errors.New("username cannot be empty")
	}

	//check to see if username is taken
	user, _ := QuerryDB(username)
	if user != nil {
		return false, errors.New("username is taken")
	}

	return true, nil

}

func validateForm(form *handlers.Form) {
	//range over the form and validate the various fields
	for field, value := range form.FieldValues {
		switch field {
		case "email":
			if isValid, err := ValidateEmail(value); !isValid {
				form.Errors[field] = err
			}
			break
		case "username":
			if isValid, err := ValidateUsername(field); !isValid {
				form.Errors[field] = err
			}
		}
	}
}
