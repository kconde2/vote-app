package models

import (
	"encoding/json"
	"errors"
)

// User represents the user.
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"pass"`
}

// Valid v
func (u User) Valid() []error {
	var errs []error
	if len(u.Password) == 0 {
		errs = append(errs, errors.New("no given password"))
	}
	if len(errs) != 0 {
		return errs
	}
	return nil
}

// MarshalJSON is marshaling the user.
func (u User) MarshalJSON() ([]byte, error) {
	type UserResponse struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var ur UserResponse
	ur.ID = u.ID
	ur.FirstName = u.FirstName
	ur.LastName = u.LastName
	return json.Marshal(ur)
}
