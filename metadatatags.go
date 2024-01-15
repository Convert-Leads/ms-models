package models

import (
    "gorm.io/gorm"
)

type MetadataTag struct {
    gorm.Model
    OrganisationId uint    `json:"organisation_id"`
    Name           string `json:"name"`
    Description    string `json:"description"`
}
