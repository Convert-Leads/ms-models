package models

import (
    "gorm.io/gorm"
)

type Media struct {
    gorm.Model
    OrganisationId uint   `json:"organisation_id"`
    Uri            string `json:"uri"`
    ContentId         uint       `json:"content_id"`
    // The reverse association is usually not serialized to JSON, but if needed:
    // Content        Content `gorm:"reference:MediaId" json:"content,omitempty"` // Reverse association to Content with JSON annotation (optional)
}