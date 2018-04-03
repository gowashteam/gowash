package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/okeyonyia123/gowash/handlers"
	"github.com/okeyonyia123/gowash/models"
)

/**
type User struct {
	UUID              string `json:"uuid" bson:"uuid"`
	Username          string `json:"username" bson:"username"`
	FirstName         string `json:"firstName" bson:"firstName"`
	LastName          string `json:"lastName" bson:"lastName"`
	Email             string `json:"email" bson:"email"`
	PasswordHash      string `json:"passwordHash" bson:"passwordHash"`
	TimestampCreated  int64  `json:"timestampCreated" bson:"timestampCreated"`
	TimestampModified int64  `json:"timestampModified" bson:"timestampModified"`
} **/

func TestHomePage(t *testing.T) {
	//step1: Create a sample Request which satisfies *http.Request
	req, err := http.NewRequest("GET", "/", nil)

	//handle error if any
	if err != nil {
		t.Fatal(err)
	}

	//step2: Create a RequestRecorder which satisfies http.ResponseWriter
	rr := httptest.NewRecorder()

	//Step3: Get a handler to make a Request
	handler := http.HandlerFunc(handlers.HomeHandler)

	//step4: Execute a request
	handler.ServeHTTP(rr, req)

	//Step5: Examine the response recorder fields for result
	if rr.Code != http.StatusOK {
		t.Errorf("Returned : %v  Expected : %v", rr.Code, http.StatusOK)
	}

}

func TestSignUp(t *testing.T) {
	//create a Request

	data := url.Values{
		"email":           {"onyia.okey@gmail.com"},
		"username":        {"kem"},
		"firstname":       {"okey"},
		"lastname":        {"onyia"},
		"password":        {"security1"},
		"confirmpassword": {"security"},
	}

	encodedData := data.Encode() //transform to a querry

	//payload := []byte(`{"name":"test product","price":11.22}`)
	//reader := strings.NewReader(payload)

	url := "http://localhost:8084/signup/?" + encodedData //form a full querry url

	req, err := http.NewRequest("POST", url, nil) //push a new request

	if err != nil {
		t.Errorf("error getting a request")
		return
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.Signup)

	handler.ServeHTTP(rr, req)

	if code := rr.Code; code != http.StatusOK {
		t.Errorf("Server returned %v instead of %v", code, http.StatusOK) // test will fail
	}

	//var promise map[string]string
	var promise *models.User
	var eroorPromise *handlers.Form
	fromServer := rr.Body.Bytes()

	json.Unmarshal(fromServer, &promise)
	json.Unmarshal(fromServer, &eroorPromise)
	// !reflect.DeepEqual(promise, payload)
	switch rr.Code == http.StatusOK {
	case true:
		if promise.Username != "dessssy" {
			t.Error("Test Failed")
			t.Error(promise)
			return
		}

	case false:
		if len(eroorPromise.Errors) < 1 {
			t.Error("Test Failed")
			t.Error(eroorPromise)
			return
		}

	default:
		t.Errorf("There is a problem with the test code")

	}

}

func TestLogin(t *testing.T) {
	//create a Request
	req, err := http.NewRequest("GET", "/login", nil)

	if err != nil {
		t.Fatal(err)
	}

	//create a response recorder
	resp := httptest.NewRecorder()

	//Get the handler to be tested
	handler := http.HandlerFunc(handlers.Login)

	//make a request
	handler.ServeHTTP(resp, req)

	//now the response recorder "resp" should have all the responses
	//go ahead and check the responses
	if code := resp.Code; code != http.StatusOK {
		t.Errorf("failed test. Status code returned %v instead of %v", code, http.StatusOK) // failed test
	}
}
