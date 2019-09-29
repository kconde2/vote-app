package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"

	"github.com/stretchr/testify/assert"
)

// Unit test that checks CreateUser method
func TestAddUser(t *testing.T) {
	
	var jsonStr = []byte(`{"first_name":"Adam","last_name":"Sow","email":"dachic@gmail.com","pass":"unit_test","birth_date":"19-10-1997"}`)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router := setupRouter()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v request %v", status, http.StatusOK, req)
	}
	router.Run(":8080")

	// Check response code
	assert.Equal(t, 200, rr.Code)
}

// Unit test that checks GetUsers method
func TestGetUsers(t *testing.T) {
	
	req, err := http.NewRequest("GET", "/users/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := setupRouter()

	// Nul pointer exception...
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v router %v",
			status, http.StatusOK, rr)
	}
	// expected := `[{""}]`
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
	assert.Equal(t, 200, rr.Code)
}
