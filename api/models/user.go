package models

import (
	"encoding/json"
	"errors"
	
	"time"
	"gopkg.in/go-playground/validator.v9"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// User represents the user.
type User struct {
	ID          int       `gorm:"primary_key"`
	UUID        uuid.UUID `json:"uuid"`
	AccessLevel int		`json:"access_level" Usage:"oneof=0 1"`
	FirstName   string `json:"first_name" validate:"required,min=2,excludesall=0x20" Usage:"alpha"`
	LastName    string `json:"last_name" validate:"required,min=2,excludesall=0x20" Usage:"alpha"`
	Email       string `json:"email" validate:"required,email" Usage:"unique"`
	Password    string `json:"pass" validate:"required" Usage:"eqfield=confirm_password"`
	// DateOfBirth time.Time `json:"birth_date" validate:"required,adult"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// Valid checks that user struct is valid
func (user User) Valid() []error {
	
	var errs []error

	v := validator.New()
	// _ = v.RegisterValidation("adult", func(fl validator.FieldLevel) bool {
	// 	now := time.Now()
	// 	if v,err := time.Parse("2019-10-19",fl.Field().String()); err == nil {
	// 		diff := now.Sub(v).Seconds()
	// 			return diff >= 18
	// 	}
	// 	return false
	// })
	if user.AccessLevel < 0 || user.AccessLevel > 1 {
		errs = append(errs, errors.New("This field's value can only be 0 or 1"))
	}

	err := v.Struct(user)
	if err != nil {
		errs = append(errs, err)
		return errs
	}

	return nil
}

// MarshalJSON is marshaling the user.
func (user User) MarshalJSON() ([]byte, error) {
	type UserResponse struct {
		ID        int    `json:"id"`
		AccessLevel int  `json:"access_level"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email  string `json:"email"`
		Password  string `json:"pass"`
		DateOfBirth time.Time `json:"birth_date"`
	}

	var ur UserResponse
	ur.ID = user.ID
	ur.FirstName = user.FirstName
	ur.LastName = user.LastName
	return json.Marshal(ur)
}

// BeforeCreate is gorm hook that is triggered before saving new user
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UUID", uuid.UUID.String)
	return nil
}

// BeforeUpdate is gorm hook that is triggered on every updated on user struct
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}
