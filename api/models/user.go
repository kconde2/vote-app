package models

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	//"github.com/kconde2/vote-app/api/utils"
)

// User represents the user.
type User struct {
	ID          int       `gorm:"primary_key"`
	UUID        uuid.UUID `json:"uuid"`
	AccessLevel int       `json:"access_level" Usage:"oneof=0 1"`
	FirstName   string    `json:"first_name" validate:"required,min=2,excludesall=0x20" Usage:"alpha"`
	LastName    string    `json:"last_name" validate:"required,min=2,excludesall=0x20" Usage:"alpha"`
	Email       string    `json:"email" validate:"required,email" Usage:"unique"`
	Password    string    `json:"pass" validate:"required" Usage:"eqfield=confirm_password"`
	DateOfBirth time.Time `json:"birth_date" validate:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// UserResponse represents user data that can be returned as response
type UserResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	DateOfBirth string    `json:"birth_date"`
}

// Valid checks that user struct is valid
func (user User) Valid() []error {

	/*var errs []error

	value := validator.New()*/
	// _ = value.RegisterValidation("adult", func(fl validator.FieldLevel) bool {
	// 	now := time.Now()
	// 	if value,err := time.Parse("2019-10-19",fl.Field().String()); err == nil {
	// 		diff := now.Sub(value).Seconds()
	// 			return diff >= 18
	// 	}
	// 	return false
	// })
	/**if user.AccessLevel < 0 || user.AccessLevel > 1 {
		errs = append(errs, errors.New("This field's value can only be 0 or 1"))
	}

	err := value.Struct(user)
	if err != nil {
		errs = append(errs, err)
		return errs
	}*/

	if age.FromDate(user.DateOfBirth) {

	}

	return nil
}

// SetPassword s
func (user *User) SetPassword(password string) {
	pwd := user.Password
	if hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost); err != nil {
		log.Println(err)
	} else {
		user.Password = string(hash)
	}
}

// BeforeCreate is gorm hook that is triggered before saving new user
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	// or error handling
	u2, err := uuid.NewV4()
	if err != nil {
		return err
	}

	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UUID", u2)
	return nil
}

// BeforeUpdate is gorm hook that is triggered on every updated on user struct
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

// MarshalJSON is marshaling the user.
func (user User) MarshalJSON() ([]byte, error) {

	var ur UserResponse
	layout := "02-01-2006"

	// Custom date in dd-mm-yyy format
	year, month, day := user.DateOfBirth.Date()
	date := strconv.Itoa(day) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(year)
	t, _ := time.Parse(layout, date)

	ur.UUID = user.UUID
	ur.FirstName = user.FirstName
	ur.LastName = user.LastName
	ur.Email = user.Email
	ur.DateOfBirth = t.Format(layout)
	return json.Marshal(ur)
}

// UnmarshalJSON iu
func (user *User) UnmarshalJSON(data []byte) error {
	var jsonData map[string]string
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}

	for key, value := range jsonData {
		if strings.ToLower(key) == "first_name" {
			user.FirstName = value
		}

		if strings.ToLower(key) == "last_name" {
			user.LastName = value
		}

		if strings.ToLower(key) == "email" {
			user.Email = value
		}

		if strings.ToLower(key) == "password" {
			user.SetPassword(value)
		}

		if strings.ToLower(key) == "birth_date" {
			dateOfBirth, err := time.Parse("02-01-2006", value)
			if err != nil {
				return err
			}
			user.DateOfBirth = dateOfBirth
		}
	}
	return nil
}
