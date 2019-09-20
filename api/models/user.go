package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User represents the user.
type User struct {
	gorm.Model
	ID          int       `gorm:"primary_key"`
	UUID        uuid.UUID `json:"uuid"`
	AccessLevel int
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"pass"`
	DateOfBirth time.Time `json:"birth_date"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// Valid v
func (user User) Valid() []error {
	var errs []error
	if len(user.Password) == 0 {
		errs = append(errs, errors.New("No given password"))
	}

	if len(errs) != 0 {
		return errs
	}

	return nil
}

// MarshalJSON is marshaling the user.
func (user User) MarshalJSON() ([]byte, error) {
	type UserResponse struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var ur UserResponse
	ur.ID = user.ID
	ur.FirstName = user.FirstName
	ur.LastName = user.LastName
	return json.Marshal(ur)
}

// BeforeCreate B
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UUID", uuid.UUID.String)
	return nil
}

// BeforeUpdate B
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

// TODO: For check for soft delete
