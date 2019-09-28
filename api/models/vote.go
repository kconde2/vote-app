package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Vote represents the vote.
type Vote struct {
	ID          int        `gorm:"primary_key"`
	UUID        uuid.UUID  `json:"uuid"`
	Title       string     `json:"title" valid:"required"`
	Description string     `json:"desc" valid:"required"`
	UUIDVote    []User     `json:"uuid_votes" gorm:"many2many:votes_users;association_foreignkey:UUID;foreignkey:uuid"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     time.Time  `json:"end_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

// VoteResponse represents vote data that can be returned as response
type VoteResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Title       string    `json:"title"`
	Description string    `json:"desc"`
}

// BeforeCreate is gorm hook that is triggered before saving new user
func (vote *Vote) BeforeCreate(scope *gorm.Scope) error {
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
func (vote *Vote) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

// MarshalJSON is marshaling the Vote.
func (vote Vote) MarshalJSON() ([]byte, error) {
	var voteResponse VoteResponse

	voteResponse.UUID = vote.UUID
	voteResponse.Title = vote.Title
	voteResponse.Description = vote.Description

	return json.Marshal(voteResponse)
}

// UnmarshalJSON create vote formatted representation of jsonData
func (vote *Vote) UnmarshalJSON(data []byte) error {
	var jsonData map[string]string
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}

	for key, value := range jsonData {
		if strings.ToLower(key) == "title" {
			vote.Title = value
		}

		if strings.ToLower(key) == "desc" {
			vote.Description = value
		}

		if strings.ToLower(key) == "start_date" {
			startDate, err := time.Parse("02-01-2006", value)
			if err != nil {
				return err
			}
			vote.StartDate = startDate
		}

		if strings.ToLower(key) == "end_date" {
			endDate, err := time.Parse("02-01-2006", value)
			if err != nil {
				return err
			}
			vote.EndDate = endDate
		}
	}
	return nil
}
