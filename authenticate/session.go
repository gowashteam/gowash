package authenticate

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/okeyonyia123/gowash/models"

	"github.com/gorilla/sessions"
)

var SessionStore *sessions.FilesystemStore

//Here is where you create a session
func CreateUserSession(u *models.User, sessionID string, w http.ResponseWriter, r *http.Request) error {
	gfSession, err := SessionStore.Get(r, "cityrescue-session")

	if err != nil {
		log.Print(err)
	}

	gfSession.Values["sessionID"] = sessionID
	gfSession.Values["username"] = u.Username
	gfSession.Values["firstName"] = u.FirstName
	gfSession.Values["lastName"] = u.LastName
	gfSession.Values["email"] = u.Email

	err = gfSession.Save(r, w)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func CreatePartnerSession(p *models.Partner, sessionID string, w http.ResponseWriter, r *http.Request) error {
	gfSession, err := SessionStore.Get(r, "gowash-session")

	if err != nil {
		log.Print(err)
	}

	gfSession.Values["sessionID"] = sessionID
	gfSession.Values["username"] = p.Username
	gfSession.Values["firstName"] = p.FirstName
	gfSession.Values["lastName"] = p.LastName
	gfSession.Values["email"] = p.Email

	err = gfSession.Save(r, w)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func init() {

	SessionStore = sessions.NewFilesystemStore("", []byte(os.Getenv("GOWASH_HASH_KEY")))
	fmt.Println([]byte(os.Getenv("GOWASH_HASH_KEY")))
	fmt.Println(os.Getenv("GOWASH_HASH_KEY"))
}
