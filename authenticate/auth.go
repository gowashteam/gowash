package authenticate

import (
	"fmt"
	"log"
	"strings"

	"github.com/okeyonyia123/gowash/util"
	"github.com/okeyonyia123/gowash/validation"
)

func VerifyCredentials(username string, password string) bool {
	db, err := validation.GetConnection()

	if err != nil {
		log.Print(err)
	} else {
		fmt.Printf("Established Connection to Database on : %v", db) //log the connection to indicate successful connection
		fmt.Println()
	}

	u, err := db.GetUser(username)
	if u == nil {
		return false
	}

	if err != nil {
		log.Print(err)
		return false
	}

	if strings.ToLower(username) == strings.ToLower(u.Username) && util.SHA256OfString(password) == u.PasswordHash {
		log.Println("Successful login attempt from user: ", u.Username)
		return true
	} else {
		log.Println("Unsuccessful login attempt from user: ", u.Username)
		return false
	}

}

func VerifyPatnerCredentials(username string, password string) bool {
	db, err := validation.GetConnection()

	if err != nil {
		log.Println("Encountered error connecting to Database : ", err)
	}

	u, err := db.GetPartner(username)
	if u == nil {
		return false
	}

	if err != nil {
		log.Print(err)
	}

	if strings.ToLower(username) == strings.ToLower(u.Username) && util.SHA256OfString(password) == u.PasswordHash {
		log.Println("Successful login attempt from user: ", u.Username)
		return true
	} else {
		log.Println("Unsuccessful login attempt from user: ", u.Username)
		return false
	}

}
