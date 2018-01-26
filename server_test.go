package developertest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/billyboar/developer-test-1/externalservice"
)

func TestPOSTCallsAndReturnsJSONfromExternalServicePOST(t *testing.T) {
	// Descirption
	//
	// Write a test that accepts a POST request on the server and sends it the
	// fake external service with the posted form body return the response.
	//
	// Use the externalservice.Client interface to create a mock and assert the
	// client was called <n> times.
	//
	// ---
	//
	// Server should receive a request on
	//
	//  [POST] /api/posts/:id
	//  application/json
	//
	// With the form body
	//
	//  application/x-www-form-urlencoded
	//	title=Hello World!
	//	description=Lorem Ipsum Dolor Sit Amen.
	//
	// The server should then relay this data to the external service by way of
	// the Client POST method and return the returned value out as JSON.
	//
	// ---
	//
	// Assert that the externalservice.Client#POST was called 1 times with the
	// provided `:id` and post body and that the returned Post (from
	// externalservice.Client#POST) is written out as `application/json`.
	server := newServer()
	go func() {
		server.Logger.Fatal(server.Start(":8080"))
	}()
	time.Sleep(time.Second) //waiting for server finish starting the server
	form := url.Values{}
	form.Add("title", "Hello World!")
	form.Add("description", "Lorem Ipsum Dolor Sit Amen.")

	req, err := http.NewRequest("POST", "http://localhost:8080/api/posts/1", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatalf("cannot create request: %s", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send POST request: %s", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var post externalservice.Post
	if err := json.Unmarshal(respBody, &post); err != nil {
		t.Errorf("invalid json data: %s", err)
	}

	if post.ID != 1 {
		t.Errorf("invalid ID. expected: 1, got: %d", post.ID)
	}
	if post.Title != "Hello World!" {
		t.Errorf("invalid title. expected 'Hello World!', got: '%s'", post.Title)
	}
	if post.Description != "Lorem Ipsum Dolor Sit Amen." {
		t.Errorf("invalid description. expected: 'Lorem Ipsum Dolor Sit Amen.', got: '%s'", post.Description)
	}
}

func TestPOSTCallsAndReturnsErrorAsJSONFromExternalServiceGET(t *testing.T) {
	// Description
	//
	// Write a test that accepts a GET request on the server and returns the
	// error returned from the external service.
	//
	// Use the externalservice.Client interface to create a mock and assert the
	// client was called <n> times.
	//
	// ---
	//
	// Server should receive a request on
	//
	//	[GET] /api/posts/:id
	//
	// The server should then return the error from the external service out as
	// JSON.
	//
	// The error response returned from the external service would look like
	//
	//	400 application/json
	//
	//	{
	//		"code": 400,
	//		"message": "Bad Request"
	//	}
	//
	// ---
	//
	// Assert that the externalservice.Client#GET was called 1 times with the
	// provided `:id` and the returned error (above) is output as the response
	// as
	//
	//	{
	//		"code": 400,
	//		"message": "Bad Request",
	//		"path": "/api/posts/:id
	//	}
	//
	// Note: *`:id` should be the actual `:id` in the original request.*
	server := newServer()
	go func() {
		server.Logger.Fatal(server.Start(":8080"))
	}()
	time.Sleep(time.Second) //waiting for server finish starting the server

	req, err := http.NewRequest("GET", "http://localhost:8080/api/posts/1", nil)
	if err != nil {
		t.Fatalf("cannot create request: %s", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send POST request: %s", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var resErr externalservice.Error
	if err := json.Unmarshal(respBody, &resErr); err != nil {
		t.Errorf("invalid json data: %s", err)
	}

	if resErr.Code != http.StatusBadRequest {
		t.Errorf("invalid error code. expected: 400, got: %d", resErr.Code)
	}

	if resErr.Message != "Bad Request" {
		t.Errorf("invalid error code. expected: 'Bad Request', got: %s", resErr.Message)
	}
}
