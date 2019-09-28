package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	// "github.com/kconde2/vote-app/api/controllers"
)

func TestAddUser(t *testing.T) {
	
	var jsonStr = []byte(`{"first_name":"Adam","last_name":"Sow","email":"dachic@gmail.com","pass":"unit_test","birth_date":"19-10-1997"}`)

	req, err := http.NewRequest("POST", "http://localhost:4000/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router := gin.Default()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v request %v", status, http.StatusOK, req)
	}

	// Check response code
	assert.Equal(t, 200, rr.Code)
}

func TestGetUsers(t *testing.T) {
	router := gin.Default()
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(controllers.GetUsers)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{""}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
