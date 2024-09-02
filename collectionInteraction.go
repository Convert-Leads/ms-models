package models

import (
	"time"
)

type CollectionInteraction struct {
	ID              uint      `json:"ID" gorm:"primaryKey"`
	OrganisationId  uint      `json:"organisation_id"`
	UserId          uint      `json:"user_id"`
	User            User      `json:"u" gorm:"foreignKey:UserId"`
	ParentType      string    `json:"parent_type"` // Should be "collection"
	ParentID        uint      `json:"parent_id"`   // ID of the collection
	InteractionType string    `json:"it"`          // "like", "comment", etc.
	InteractionData string    `json:"ida"`         // Additional data if needed (e.g., comment text)
	CreatedAt       time.Time `json:"ca"`
	UpdatedAt       time.Time `json:"ua"`
}
