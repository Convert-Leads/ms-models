package models

import (
	"time"
)

type OrganisationAPIKey struct {
	Model
	OrganisationID uint        `gorm:"not null" json:"organisation_id"`
	APIKey         string       `gorm:"type:varchar(255);unique;not null" json:"api_key"`
	AccessLevel           string       `gorm:"type:varchar(50);not null" json:"access_level"`
	Description    string       `gorm:"type:text" json:"description,omitempty"`
	LastUsed       *time.Time   `json:"last_used,omitempty"`
	ExpiresAt      *time.Time   `json:"expires_at,omitempty"`
	CreatedBy      uint         `json:"created_by,omitempty"`
}

