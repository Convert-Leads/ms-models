package models

import (
    "gorm.io/gorm"
)

type ContentType struct {
    gorm.Model
    OrganisationId int    `json:"organisation_id"`
    Name           string `json:"name"`
    Description    string `json:"description"`
}
