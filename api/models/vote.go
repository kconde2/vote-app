package models

import (
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Vote represents the vote.
type Vote struct {
	ID          int        `gorm:"primary_key"`
	UUID        uuid.UUID  `json:"uuid"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	UUIDVote    []User     `json:"uuid_vote" gorm:"many2many:votes_users;association_foreignkey:UUID;foreignkey:uuid"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     time.Time  `json:"end_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
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
		UUID        uuid.UUID `json:"uuid"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
	}

	var vr VoteResponse
	vr.UUID = u.UUID
	vr.Title = u.Title
	return json.Marshal(vr)
}
