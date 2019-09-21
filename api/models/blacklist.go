package models

import (
	"time"
)

// Blacklist represents blacklist's IP that aren't allowed to login
type Blacklist struct {
	ID        int    `gorm:"primary_key"`
	IPAddress string `json:"ip_address" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
