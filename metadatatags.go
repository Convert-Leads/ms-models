package models

import (
    "gorm.io/gorm"
)

type MetadataTag struct {
    gorm.Model
    OrganisationId uint    `json:"o"`
    Name           string `json:"n"`
    Description    string `json:"d,omitempty"`
}
