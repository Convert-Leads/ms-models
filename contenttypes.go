package models

import (
    "gorm.io/gorm"
)

type ContentTypes struct {
    gorm.Model
    OrganisationId int    `json:"organisation_id"`
    Name           string `json:"name"`
    Description    string `json:"description"`
}
