package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/okeyonyia123/gowash/servers/Signup/handlers"
)

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

	req, err := http.NewRequest("GET", "http://localhost:8084/signup", nil)

	//handle the error
	if err != nil {
		fmt.Println("error creating a request")
		t.Fatal(err) // This will fail the test
	}
	// create a Client to run a Request
	client := http.Client{}

	//create a request with the client and receive a response back
	resp, err := client.Do(req)
	//handle the error
	if err != nil {
		t.Fatal(err) // This will fail the test
	}

	// Now at this point we have a response and we can read all it says
	//lets close the response body last
	defer resp.Body.Close()

	//Lets read the status code and make sure it returns OK
	if code := resp.StatusCode; code != 200 {
		t.Errorf("Test failed, server returned a status code %v instead of 200", code) // fail the test
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
