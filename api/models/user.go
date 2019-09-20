package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type myTime time.Time

var _ json.Unmarshaler = &myTime{}

func (mt *myTime) UnmarshalJSON(bs []byte) error {
	var s map[string]string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}

	t, err := time.Parse("2008-19-12", "2008-19-12")
	if err != nil {
		return err
	}

	*mt = myTime(t)
	return nil
}

// User represents the user.
type User struct {
	ID          int       `gorm:"primary_key"`
	UUID        uuid.UUID `json:"uuid"`
	AccessLevel int
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"pass"`
	DateOfBirth myTime `json:"birth_date"`
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
