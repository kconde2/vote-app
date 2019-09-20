package models

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
)

// Vote represents the vote.
type Vote struct {
	gorm.Model
	ID          int        `json:"id"`
	UUID        string     `json:"uuid"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	UUIDVote    []*string  `json:"uuid_vote"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     time.Time  `json:"end_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeleteAt    *time.Time `json:"delete_at"`
}

// Valid v
func (u Vote) Valid() []error {
	var errs []error

	if len(errs) != 0 {
		return errs
	}
	return nil
}

// MarshalJSON is marshaling the Vote.
func (u Vote) MarshalJSON() ([]byte, error) {
	type VoteResponse struct {
		UUID        string `json:"uuid"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	var vr VoteResponse
	vr.UUID = u.UUID
	vr.Title = u.Title
	return json.Marshal(vr)
}
