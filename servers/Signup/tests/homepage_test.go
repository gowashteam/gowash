package tests

import (
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
