package models

import (
	"time"
)

// Blacklist b.
type Blacklist struct {
	ID        int    `gorm:"primary_key"`
	IPAddress string `json:"ip_address" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
